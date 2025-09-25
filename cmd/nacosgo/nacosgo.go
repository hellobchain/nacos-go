package nacosgo

import (
	"fmt"
	"net/http"
	"time"

	"github.com/hellobchain/nacos-go/conf"
	"github.com/hellobchain/nacos-go/config"
	"github.com/hellobchain/nacos-go/constant"
	"github.com/hellobchain/nacos-go/handle"
	"github.com/hellobchain/nacos-go/middleware"
	"github.com/hellobchain/nacos-go/service"
	"github.com/hellobchain/nacos-go/tenant"
	"github.com/hellobchain/nacos-go/user"
	"github.com/hellobchain/wswlog/wlogging"
)

var logger = wlogging.MustGetFileLoggerWithoutName(nil)

func StartServer(allService conf.AllService, serverPort int) {
	// 启动 HTTP：把两个路由挂载到同端口
	r := handle.NewLogRouter()
	// 全局中间件（顺序：CORS → Logger → JWTAuth）
	r.Use(middleware.CORS, middleware.Logger, middleware.JWTAuth)
	service.RegistryRoute(r, allService.InstanceService) // 原 /nacos/v1/ns 路由
	config.ConfigRoute(r, allService.ConfigService)      // 新增 /v1/cs 路由
	tenant.TetantRoute(r, allService.TenantService)      // 新增 /v1/tenant
	user.AuthRoute(r, allService.UserService)            // 新增 /v1/auth
	if serverPort == 0 {
		logger.Warnf("Invalid server port %d, use default %d", serverPort, constant.DEFAULT_SERVER_PORT)
		serverPort = constant.DEFAULT_SERVER_PORT
	}
	addr := fmt.Sprintf(":%d", serverPort)
	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}
	logger.Infof("nacos-go listen %s", addr)
	logger.Fatal(srv.ListenAndServe())
}

// startHeartbeat 每 5 秒扫描一次过期实例
func StartHeartbeat(svc *service.RegistryService, heartBeatInternal time.Duration) {
	if heartBeatInternal <= 0 {
		logger.Warnf("Invalid heartbeat interval %d, use default 5", heartBeatInternal)
		heartBeatInternal = 5
	}
	go func() {
		ticker := time.NewTicker(heartBeatInternal * time.Second)
		defer ticker.Stop()
		for range ticker.C {
			logger.Debug("Clean expired instances")
			// 如果底层 repo 实现了 CleanExpired，直接调
			if cleaner, ok := svc.Repo.(interface{ CleanExpired() error }); ok {
				_ = cleaner.CleanExpired()
			}
		}
	}()
}

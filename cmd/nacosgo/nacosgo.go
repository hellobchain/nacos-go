package nacosgo

import (
	"net/http"
	"time"

	"github.com/hellobchain/nacos-go/config"
	"github.com/hellobchain/nacos-go/handle"
	"github.com/hellobchain/nacos-go/service"
	"github.com/hellobchain/wswlog/wlogging"
)

var logger = wlogging.MustGetFileLoggerWithoutName(nil)

func StartServer(regSvc *service.RegistryService, confSvc *config.Service) {
	// 启动 HTTP：把两个路由挂载到同端口
	r := handle.NewLogRouter()
	service.RegistryRoute(r, regSvc) // 原 /nacos/v1/ns 路由
	config.ConfigRoute(r, confSvc)   // 新增 /v1/cs 路由
	srv := &http.Server{
		Addr:         ":8848",
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	logger.Info("nacos-go listen :8848")
	logger.Fatal(srv.ListenAndServe())
}

// startHeartbeat 每 5 秒扫描一次过期实例
func StartHeartbeat(svc *service.RegistryService) {
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()
		for range ticker.C {
			// 如果底层 repo 实现了 CleanExpired，直接调
			if cleaner, ok := svc.Repo.(interface{ CleanExpired() error }); ok {
				_ = cleaner.CleanExpired()
			}
		}
	}()
}

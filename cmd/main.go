package main

import (
	"flag"

	"github.com/hellobchain/nacos-go/cmd/nacosgo"
	"github.com/hellobchain/nacos-go/conf"
	"github.com/hellobchain/nacos-go/config"
	"github.com/hellobchain/nacos-go/service"
	"github.com/hellobchain/wswlog/wlogging"
)

var logger = wlogging.MustGetFileLoggerWithoutName(nil)
var (
	driver = flag.String("driver", "memory", "memory/mysql/postgres")
	dsn    = flag.String("dsn", "", "sql dsn")
)

func main() {
	wlogging.SetConsole(true)
	wlogging.SetGlobalLogLevel("debug")
	flag.Parse()
	// 1. 初始化 DAO（mysql / memory / mysql / postgres)
	repo, db, err := conf.Init(*driver, *dsn)
	if err != nil {
		logger.Fatal("init error: %v", err)
	}
	// 2. 初始化配置服务
	configRepo := conf.InitConfig(*driver, db)
	// 3. 构建业务服务
	regSvc := service.NewRegistryService(repo)
	// 4. 构建业务服务
	confSvc := config.NewService(configRepo)
	// 5. 启动后台心跳清理协程
	nacosgo.StartHeartbeat(regSvc)
	// 6. 启动 HTTP 服务（阻塞）
	nacosgo.StartServer(regSvc, confSvc)
}

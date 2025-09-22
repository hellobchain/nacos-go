package main

import (
	"flag"

	"github.com/hellobchain/nacos-go/cmd/nacosgo"
	"github.com/hellobchain/nacos-go/conf"
	daoConfig "github.com/hellobchain/nacos-go/config"
	"github.com/hellobchain/nacos-go/pkg/config"
	"github.com/hellobchain/nacos-go/service"
	"github.com/hellobchain/wswlog/wlogging"
)

var logger = wlogging.MustGetFileLoggerWithoutName(nil)
var (
	confPath = flag.String("conf", "nacos.yml", "config path")
)

func main() {
	flag.Parse()
	nacosConfig, err := config.InitNacosConfig(*confPath)
	if err != nil {
		logger.Fatal("init nacos config error: %v", err)
		return
	}
	// 0. init log
	wlogging.SetConsole(nacosConfig.ServerConfig.Console)
	wlogging.SetGlobalLogLevel(nacosConfig.ServerConfig.LogLevel)
	// 1. 初始化 DAO（mysql / memory / mysql / postgres)
	repo, db, err := conf.Init(nacosConfig.DBConfig.Driver, nacosConfig.DBConfig.Dsn)
	if err != nil {
		logger.Fatal("init error: %v", err)
		return
	}
	// 2. 初始化配置服务
	configRepo := conf.InitConfig(nacosConfig.DBConfig.Driver, db)
	// 3. 构建业务服务
	regSvc := service.NewRegistryService(repo)
	// 4. 构建业务服务
	confSvc := daoConfig.NewService(configRepo)
	// 5. 启动后台心跳清理协程
	nacosgo.StartHeartbeat(regSvc, nacosConfig.ServerConfig.HeartbeatInterval)
	// 6. 启动 HTTP 服务（阻塞）
	nacosgo.StartServer(regSvc, confSvc, nacosConfig.ServerConfig.Port)
}

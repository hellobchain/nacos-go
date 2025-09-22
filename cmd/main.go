package main

import (
	"flag"

	"github.com/hellobchain/nacos-go/cmd/nacosgo"
	"github.com/hellobchain/nacos-go/conf"
	"github.com/hellobchain/nacos-go/pkg/config"
)

var (
	confPath = flag.String("conf", "nacos.yml", "config path")
)

func main() {
	flag.Parse()
	nacosConfig := config.InitNacosConfig(*confPath)
	allConfig := conf.InitAllConfig(nacosConfig.DBConfig.Driver, nacosConfig.DBConfig.Dsn)
	allService := conf.InitAllService(allConfig)
	nacosgo.StartHeartbeat(allService.InstanceService, nacosConfig.ServerConfig.HeartbeatInterval)
	nacosgo.StartServer(allService, nacosConfig.ServerConfig.Port)
}

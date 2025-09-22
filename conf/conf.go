package conf

import (
	"github.com/hellobchain/nacos-go/config"
	configMemory "github.com/hellobchain/nacos-go/config/memory"
	configMysql "github.com/hellobchain/nacos-go/config/mysql"
	"github.com/hellobchain/nacos-go/service"
	daoMemory "github.com/hellobchain/nacos-go/service/memory"
	daoMysql "github.com/hellobchain/nacos-go/service/mysql"
	"github.com/hellobchain/nacos-go/user"
	userMemory "github.com/hellobchain/nacos-go/user/memory"
	userMysql "github.com/hellobchain/nacos-go/user/mysql"
	"github.com/hellobchain/wswlog/wlogging"
)

var logger = wlogging.MustGetFileLoggerWithoutName(nil)

type AllConfig struct {
	InstanceRepo service.InstanceRepo
	UserRepo     user.UserRepo
	ConfigRepo   config.ConfigRepo
}

func InitAllConfig(driver string, dsn string) AllConfig {
	var allConfig AllConfig
	switch driver {
	case "mysql":
		logger.Debug("init service mysql")
		dbService, db := daoMysql.New(dsn)
		allConfig.ConfigRepo = configMysql.NewConfigRepo(db)
		allConfig.InstanceRepo = dbService
		allConfig.UserRepo = userMysql.New(db)
	case "memory":
		fallthrough
	default:
		logger.Debug("init service memory")
		allConfig.ConfigRepo = configMemory.New()
		allConfig.InstanceRepo = daoMemory.New()
		allConfig.UserRepo = userMemory.New()
	}
	return allConfig
}

type AllService struct {
	InstanceService *service.RegistryService
	ConfigService   *config.Service
	UserService     *user.AuthUserService
}

func InitAllService(allConfig AllConfig) AllService {
	var allService AllService
	allService.InstanceService = service.NewRegistryService(allConfig.InstanceRepo)
	allService.ConfigService = config.NewService(allConfig.ConfigRepo)
	allService.UserService = user.NewAuthUserService(allConfig.UserRepo)
	user.InitAdminUser(allService.UserService)
	return allService
}

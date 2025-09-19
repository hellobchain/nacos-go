package conf

import (
	"github.com/hellobchain/nacos-go/config"
	"github.com/hellobchain/nacos-go/config/memory"
	"github.com/hellobchain/nacos-go/config/mysql"
	"github.com/hellobchain/nacos-go/service"
	daoMemory "github.com/hellobchain/nacos-go/service/memory"
	daoMysql "github.com/hellobchain/nacos-go/service/mysql"
	"github.com/hellobchain/wswlog/wlogging"

	"gorm.io/gorm"
)

var logger = wlogging.MustGetFileLoggerWithoutName(nil)

func Init(driver, dsn string) (service.InstanceRepo, *gorm.DB, error) {
	switch driver {
	case "mysql":
		logger.Debug("init service mysql")
		return daoMysql.New(dsn)
	case "memory":
		fallthrough
	default:
		logger.Debug("init service memory")
		return daoMemory.New(), nil, nil
	}
}

// InitConfig 根据 driver 初始化配置中心存储
func InitConfig(driver string, db interface{}) config.ConfigRepo {
	switch driver {
	case "mysql":
		logger.Debug("init config mysql")
		return mysql.NewConfigRepo(db.(*gorm.DB))
	case "memory":
		fallthrough
	default:
		logger.Debug("init config memory")
		return memory.New()
	}
}

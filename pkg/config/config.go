package config

import (
	"strings"
	"time"

	"github.com/hellobchain/wswlog/wlogging"
	"github.com/hokaccha/go-prettyjson"
	"github.com/spf13/viper"
)

var logger = wlogging.MustGetFileLoggerWithoutName(nil)

type NacosConfig struct {
	ServerConfig ServerConfig `mapstructure:"server"` // 服务器配置
	DBConfig     DBConfig     `mapstructure:"db"`     // 数据库配置
}

type ServerConfig struct {
	Port              int           `mapstructure:"port"`               // 端口
	LogLevel          string        `mapstructure:"log_level"`          // 日志级别
	Console           bool          `mapstructure:"console"`            // 是否开启控制台输出
	HeartbeatInterval time.Duration `mapstructure:"heartbeat_interval"` // 心跳检测间隔
}

// 数据库类型配置
type DBConfig struct {
	Driver string `mapstructure:"driver"` // 数据库驱动
	Dsn    string `mapstructure:"dsn"`    // 数据库连接串
}

const (
	// pre
	cmdPre = "NACOS"
)

func setEnvVariables() {
	// For environment variables.
	viper.SetEnvPrefix(cmdPre)
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
}

var GlobalNacosConfig NacosConfig

func InitNacosConfig(path string) (*NacosConfig, error) {
	if path == "" {
		logger.Info("use default config path")
		path = "nacos.yml"
	}
	setEnvVariables()
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		logger.Error("init nacos config error: %v", err)
		return nil, err
	}
	err := viper.Unmarshal(&GlobalNacosConfig)
	if err != nil {
		logger.Error("unmarshal nacos config error: %v", err)
		return nil, err
	}
	GlobalNacosConfig.printLog()
	return &GlobalNacosConfig, nil
}

func (c *NacosConfig) printLog() {
	json, err := prettyjson.Marshal(c)
	if err != nil {
		logger.Fatalf("marshal alarm config failed, %s", err.Error())
	}
	logger.Debug(string(json))
}

package mysql

import (
	"context"
	"encoding/json"
	"time"

	"github.com/hellobchain/nacos-go/model"
	"github.com/hellobchain/nacos-go/service"
	"github.com/hellobchain/wswlog/wlogging"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

var logger = wlogging.MustGetFileLoggerWithoutName(nil)
var _ service.InstanceRepo = (*mysqlRepo)(nil)

type sqlLogger struct {
	logger *wlogging.WswLogger
}

// LogMode(LogLevel) Interface
func (s *sqlLogger) LogMode(level gormLogger.LogLevel) gormLogger.Interface {
	return NewSqlLogger(logger)
}

// Info(context.Context, string, ...interface{})
func (s *sqlLogger) Info(ctx context.Context, str string, args ...interface{}) {
	s.logger.Infof(str, args...)
}

// Warn(context.Context, string, ...interface{})
func (s *sqlLogger) Warn(ctx context.Context, str string, args ...interface{}) {
	s.logger.Warnf(str, args...)
}

// Error(context.Context, string, ...interface{})
func (s *sqlLogger) Error(ctx context.Context, str string, args ...interface{}) {
	s.logger.Errorf(str, args...)
}

// Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error)
func (s *sqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	sql, rows := fc()
	if err != nil {
		s.logger.Errorf("[SQL] [ERROR] %s [%.3fms] [rows:%d] %v", sql, float64(time.Since(begin).Nanoseconds())/1e6, rows, err)
	}
}

func NewSqlLogger(logger *wlogging.WswLogger) *sqlLogger {
	return &sqlLogger{logger: logger}
}

type InstancePO struct {
	gorm.Model
	ServiceName string  `gorm:"column:service_name;size:128;not null;index:uk_instance,unique"`
	GroupName   string  `gorm:"column:group_name;size:128;not null;index:uk_instance,unique"`
	ClusterName string  `gorm:"column:cluster_name;size:128;default:DEFAULT"`
	IP          string  `gorm:"column:ip;size:64;not null;index:uk_instance,unique"`
	Port        uint64  `gorm:"column:port;not null;index:uk_instance,unique"`
	Weight      float64 `gorm:"column:weight;default:1"`
	Healthy     bool    `gorm:"column:healthy;default:1"`
	Ephemeral   bool    `gorm:"column:ephemeral;default:1"`
	Metadata    string  `gorm:"column:metadata;type:text"`
	ExpireTime  int64   `gorm:"column:expire_time;default:0"`
}

func (InstancePO) TableName() string { return "instance" }

type mysqlRepo struct {
	db *gorm.DB
}

// New 创建 MySQL 实现，dsn 样例：user:pass@tcp(127.0.0.1:3306)/nacos?charset=utf8mb4&parseTime=True&loc=Local
func New(dsn string) (service.InstanceRepo, *gorm.DB) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: NewSqlLogger(logger),
	})
	if err != nil {
		logger.Fatal("mysql connect error:", err)
	}
	err = db.AutoMigrate(&InstancePO{})
	if err != nil {
		logger.Fatal("mysql auto migrate service error:", err)
	}
	return &mysqlRepo{db: db}, db
}

func (r *mysqlRepo) Register(ins model.Instance) error {
	po := toPO(ins)
	return r.db.Where("service_name=? AND group_name=? AND ip=? AND port=?",
		po.ServiceName, po.GroupName, po.IP, po.Port).
		Assign(po).
		FirstOrCreate(&po).Error
}

func (r *mysqlRepo) Deregister(serviceName, groupName, ip string, port uint64) error {
	return r.db.Where("service_name=? AND group_name=? AND ip=? AND port=?",
		serviceName, groupName, ip, port).
		Delete(&InstancePO{}).Error
}

func (r *mysqlRepo) List(serviceName, groupName string) ([]model.Instance, error) {
	var pos []InstancePO
	err := r.db.Where("service_name=? AND group_name=? AND healthy=1", serviceName, groupName).Find(&pos).Error
	return toDTOList(pos), err
}

func (r *mysqlRepo) Heartbeat(serviceName, groupName, ip string, port uint64) error {
	return r.db.Model(&InstancePO{}).
		Where("service_name=? AND group_name=? AND ip=? AND port=?", serviceName, groupName, ip, port).
		Update("expire_time", time.Now().Unix()+30).Error
}

func (r *mysqlRepo) CleanExpired() error {
	return r.db.Where("expire_time>0 AND expire_time<?", time.Now().Unix()).Delete(&InstancePO{}).Error
}

/* ---------- 转换 ---------- */
func toPO(m model.Instance) InstancePO {
	meta, _ := json.Marshal(m.Metadata)
	return InstancePO{
		ServiceName: m.ServiceName,
		GroupName:   m.GroupName,
		ClusterName: m.ClusterName,
		IP:          m.IP,
		Port:        m.Port,
		Weight:      m.Weight,
		Healthy:     m.Healthy,
		Ephemeral:   m.Ephemeral,
		Metadata:    string(meta),
		ExpireTime:  m.ExpireTime,
	}
}

func toDTOList(pos []InstancePO) []model.Instance {
	res := make([]model.Instance, 0, len(pos))
	for _, p := range pos {
		res = append(res, model.Instance{
			ID:          uint64(p.ID),
			ServiceName: p.ServiceName,
			GroupName:   p.GroupName,
			ClusterName: p.ClusterName,
			IP:          p.IP,
			Port:        p.Port,
			Weight:      p.Weight,
			Healthy:     p.Healthy,
			Ephemeral:   p.Ephemeral,
			Metadata:    p.Metadata, // 兼容以前的设计，metadata 直接是字符串
			ExpireTime:  p.ExpireTime,
		})
	}
	return res
}

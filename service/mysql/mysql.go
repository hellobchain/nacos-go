package mysql

import (
	"encoding/json"
	"time"

	"github.com/hellobchain/nacos-go/model"
	"github.com/hellobchain/nacos-go/service"
	"github.com/hellobchain/wswlog/wlogging"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var logger = wlogging.MustGetFileLoggerWithoutName(nil)
var _ service.InstanceRepo = (*mysqlRepo)(nil)

type mysqlRepo struct {
	db *gorm.DB
}

// New 创建 MySQL 实现，dsn 样例：user:pass@tcp(127.0.0.1:3306)/nacos?charset=utf8mb4&parseTime=True&loc=Local
func New(dsn string) (service.InstanceRepo, *gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}
	err = db.AutoMigrate(&InstancePO{})
	if err != nil {
		logger.Fatal("mysql auto migrate error:", err)
	}
	return &mysqlRepo{db: db}, db, nil
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

package mysql

import "gorm.io/gorm"

type InstancePO struct {
	gorm.Model
	ServiceName string  `gorm:"column:service_name;index:idx_sg"`
	GroupName   string  `gorm:"column:group_name;index:idx_sg"`
	ClusterName string  `gorm:"column:cluster_name"`
	IP          string  `gorm:"column:ip"`
	Port        uint64  `gorm:"column:port"`
	Weight      float64 `gorm:"column:weight"`
	Healthy     bool    `gorm:"column:healthy"`
	Ephemeral   bool    `gorm:"column:ephemeral"`
	Metadata    string  `gorm:"column:metadata"`
	ExpireTime  int64   `gorm:"column:expire_time;index:idx_expire"`
}

func (InstancePO) TableName() string { return "instance" }

package mysql

import (
	"gorm.io/gorm"
)

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

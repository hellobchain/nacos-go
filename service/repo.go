package service

import "github.com/hellobchain/nacos-go/model"

type InstanceRepo interface {
	Register(ins model.Instance) error
	Deregister(serviceName, groupName, ip string, port uint64) error
	List(serviceName, groupName string) ([]model.Instance, error)
	Heartbeat(serviceName, groupName, ip string, port uint64) error
	CleanExpired() error // 可定时删过期
}

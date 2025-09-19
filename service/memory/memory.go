package memory

import (
	"sync"
	"time"

	"github.com/hellobchain/nacos-go/model"
	"github.com/hellobchain/nacos-go/service"
)

var _ service.InstanceRepo = (*memoryRepo)(nil)

type memoryRepo struct {
	table map[string][]model.Instance
	mu    sync.RWMutex
}

func New() service.InstanceRepo {
	table := make(map[string][]model.Instance)
	return &memoryRepo{
		table: table,
	}
}

// 注册
func (m *memoryRepo) Register(ins model.Instance) error {
	key := ins.ServiceName + "@@" + ins.GroupName
	m.mu.Lock()
	defer m.mu.Unlock()
	// 简单去重：同 ip+port 覆盖
	for i, old := range m.table[key] {
		if old.IP == ins.IP && old.Port == ins.Port {
			m.table[key][i] = ins
			return nil
		}
	}
	m.table[key] = append(m.table[key], ins)
	return nil
}

// 发现
func (m *memoryRepo) List(serviceName, groupName string) ([]model.Instance, error) {
	key := serviceName + "@@" + groupName
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.table[key], nil
}

// 心跳保活（简易版：每 5s 打一次日志）
func (m *memoryRepo) Heartbeat(serviceName, groupName, ip string, port uint64) error {
	key := serviceName + "@@" + groupName
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, ins := range m.table[key] {
		if ins.IP == ip && ins.Port == port {
			// 更新
			m.table[key][i].ExpireTime = time.Now().Unix() + 30
			break
		}
	}
	return nil
}

// 注销
func (m *memoryRepo) Deregister(serviceName, groupName, ip string, port uint64) error {
	key := serviceName + "@@" + groupName
	m.mu.Lock()
	defer m.mu.Unlock()
	instances := m.table[key]
	for i, ins := range instances {
		if ins.IP == ip && ins.Port == port {
			// 删除
			m.table[key] = append(instances[:i], instances[i+1:]...)
			break
		}
	}
	return nil
}

// 清理过期实例（简易版：不做任何事）
func (m *memoryRepo) CleanExpired() error {
	return nil
}

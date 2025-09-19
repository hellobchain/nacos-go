package service

import (
	"github.com/hellobchain/nacos-go/model"
)

type RegistryService struct {
	Repo InstanceRepo
}

func NewRegistryService(r InstanceRepo) *RegistryService {
	return &RegistryService{Repo: r}
}

func (s *RegistryService) Register(ins model.Instance) error {
	return s.Repo.Register(ins)
}

func (s *RegistryService) List(serviceName, groupName string) ([]model.Instance, error) {
	return s.Repo.List(serviceName, groupName)
}

func (s *RegistryService) Heartbeat(serviceName, groupName, ip string, port uint64) error {
	return s.Repo.Heartbeat(serviceName, groupName, ip, port)
}

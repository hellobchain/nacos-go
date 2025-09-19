package config

import (
	"context"
)

type Service struct {
	repo ConfigRepo
}

func NewService(r ConfigRepo) *Service {
	return &Service{repo: r}
}

func (s *Service) Publish(ctx context.Context, item ConfigItem) error {
	return s.repo.Save(ctx, item)
}

func (s *Service) Get(ctx context.Context, dataId, group, tenant string) (*ConfigItem, error) {
	return s.repo.Get(ctx, dataId, group, tenant)
}

func (s *Service) Delete(ctx context.Context, dataId, group, tenant string) error {
	return s.repo.Delete(ctx, dataId, group, tenant)
}

func (s *Service) List(ctx context.Context, dataId, group, tenant string) ([]ConfigItem, error) {
	return s.repo.List(ctx, dataId, group, tenant)
}

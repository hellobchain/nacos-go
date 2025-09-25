package tenant

import "context"

type Service struct {
	repo TenantRepo
}

func NewService(r TenantRepo) *Service { return &Service{repo: r} }

func (s *Service) Create(ctx context.Context, name string) error { return s.repo.Save(ctx, name) }
func (s *Service) Delete(ctx context.Context, name string) error { return s.repo.Delete(ctx, name) }
func (s *Service) List(ctx context.Context) ([]string, error)    { return s.repo.List(ctx) }

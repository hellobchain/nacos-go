package tenant

import "context"

type TenantRepo interface {
	Save(ctx context.Context, name string) error
	Delete(ctx context.Context, name string) error
	List(ctx context.Context) ([]string, error)
}

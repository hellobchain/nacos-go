package config

import "context"

type ConfigRepo interface {
	Save(ctx context.Context, item ConfigItem) error
	Get(ctx context.Context, dataId, group, tenant string) (*ConfigItem, error)
	Delete(ctx context.Context, dataId, group, tenant string) error
	List(ctx context.Context, dataId, group, tenant string) ([]ConfigItem, error)
	Edit(ctx context.Context, dataId, group, tenant, content string) error
}

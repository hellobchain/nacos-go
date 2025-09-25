package mysql

import (
	"context"
	"time"

	"github.com/hellobchain/nacos-go/tenant"
	"github.com/hellobchain/wswlog/wlogging"
	"gorm.io/gorm"
)

var logger = wlogging.MustGetFileLoggerWithoutName(nil)
var _ tenant.TenantRepo = (*tenantMysqlRepo)(nil)

type tenantMysqlRepo struct{ db *gorm.DB }
type Model struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func New(db *gorm.DB) tenant.TenantRepo {
	err := db.AutoMigrate(&tenantPO{})
	if err != nil {
		logger.Fatal("mysql auto migrate tenant error:", err)
	}
	return &tenantMysqlRepo{db: db}
}

type tenantPO struct {
	Model
	Name string `gorm:"size:128;unique;not null"`
}

func (tenantPO) TableName() string { return "tenant" }

func (r *tenantMysqlRepo) Save(ctx context.Context, name string) error {
	return r.db.WithContext(ctx).Where("name = ?", name).
		FirstOrCreate(&tenantPO{Name: name}).Error
}

func (r *tenantMysqlRepo) Delete(ctx context.Context, name string) error {
	return r.db.WithContext(ctx).Where("name = ?", name).Delete(&tenantPO{}).Error
}

func (r *tenantMysqlRepo) List(ctx context.Context) ([]string, error) {
	var names []string
	err := r.db.WithContext(ctx).Model(&tenantPO{}).Pluck("name", &names).Error
	return names, err
}

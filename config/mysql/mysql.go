package mysql

import (
	"context"
	"crypto/md5"
	"encoding/hex"

	"github.com/hellobchain/nacos-go/config"
	"github.com/hellobchain/wswlog/wlogging"
	"gorm.io/gorm"
)

var logger = wlogging.MustGetFileLoggerWithoutName(nil)

type configMysql struct {
	db *gorm.DB
}

func NewConfigRepo(db *gorm.DB) config.ConfigRepo {
	err := db.AutoMigrate(&configPO{})
	if err != nil {
		logger.Fatal("mysql auto migrate error:", err)
	}
	return &configMysql{db: db}
}

type configPO struct {
	gorm.Model
	DataId   string `gorm:"column:data_id;size:128;not null;index:uk_config,unique"`
	GroupId  string `gorm:"column:group_id;size:128;not null;default:DEFAULT_GROUP;index:uk_config,unique"`
	TenantId string `gorm:"column:tenant_id;size:128;default:'';index:uk_config,unique"`
	Content  string `gorm:"column:content;type:text;not null"`
	Md5      string `gorm:"column:md5;size:32;not null"`
	BetaIps  string `gorm:"column:beta_ips;size:1024"`
	SrcUser  string `gorm:"column:src_user;size:128"`
	SrcIp    string `gorm:"column:src_ip;size:64"`
	AppName  string `gorm:"column:app_name;size:128"`
	Type     string `gorm:"column:type;size:16;default:yaml"`
}

func (configPO) TableName() string { return "config" }

func (r *configMysql) Save(ctx context.Context, item config.ConfigItem) error {
	if item.Md5 == "" {
		item.Md5 = md5Str(item.Content)
	}
	po := configPO{
		DataId:   item.DataId,
		GroupId:  item.Group,
		TenantId: item.Tenant,
		Content:  item.Content,
		Md5:      item.Md5,
		BetaIps:  item.BetaIps,
		SrcUser:  item.SrcUser,
		SrcIp:    item.SrcIp,
		AppName:  item.AppName,
		Type:     item.Type,
	}
	return r.db.WithContext(ctx).Where("data_id=? AND group_id=? AND tenant_id=?", item.DataId, item.Group, item.Tenant).
		Assign(po).
		FirstOrCreate(&po).Error
}

func (r *configMysql) Get(ctx context.Context, dataId, group, tenant string) (*config.ConfigItem, error) {
	var po configPO
	err := r.db.WithContext(ctx).Where("data_id=? AND group_id=? AND tenant_id=?", dataId, group, tenant).First(&po).Error
	if err != nil {
		return nil, err
	}
	return &config.ConfigItem{
		DataId:  po.DataId,
		Group:   po.GroupId,
		Tenant:  po.TenantId,
		Content: po.Content,
		Md5:     po.Md5,
		Type:    po.Type,
	}, nil
}

func (r *configMysql) Delete(ctx context.Context, dataId, group, tenant string) error {
	return r.db.WithContext(ctx).Where("data_id=? AND group_id=? AND tenant_id=?", dataId, group, tenant).Delete(&configPO{}).Error
}

func (r *configMysql) List(ctx context.Context, dataId, group, tenant string) ([]config.ConfigItem, error) {
	var pos []configPO
	db := r.db.WithContext(ctx)
	if dataId != "" {
		db = db.Where("data_id LIKE ?", "%"+dataId+"%")
	}
	if group != "" {
		db = db.Where("group_id = ?", group)
	}
	if tenant != "" {
		db = db.Where("tenant_id = ?", tenant)
	}
	err := db.Find(&pos).Error
	if err != nil {
		return nil, err
	}
	items := make([]config.ConfigItem, 0, len(pos))
	for _, p := range pos {
		items = append(items, config.ConfigItem{
			DataId:  p.DataId,
			Group:   p.GroupId,
			Tenant:  p.TenantId,
			Content: p.Content,
			Md5:     p.Md5,
			Type:    p.Type,
		})
	}
	return items, nil
}

func md5Str(s string) string {
	h := md5.Sum([]byte(s))
	return hex.EncodeToString(h[:])
}

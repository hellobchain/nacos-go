package mysql

import (
	"context"
	"crypto/md5"
	"encoding/hex"

	"github.com/hellobchain/nacos-go/config"
	"gorm.io/gorm"
)

type configMysql struct {
	db *gorm.DB
}

func NewConfigRepo(db *gorm.DB) config.ConfigRepo {
	return &configMysql{db: db}
}

type configPO struct {
	gorm.Model
	DataId   string `gorm:"column:data_id;index:uk,unique"`
	GroupId  string `gorm:"column:group_id;index:uk,unique"`
	TenantId string `gorm:"column:tenant_id;index:uk,unique"`
	Content  string `gorm:"column:content"`
	Md5      string `gorm:"column:md5"`
	BetaIps  string `gorm:"column:beta_ips"`
	SrcUser  string `gorm:"column:src_user"`
	SrcIp    string `gorm:"column:src_ip"`
	AppName  string `gorm:"column:app_name"`
	Type     string `gorm:"column:type"`
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

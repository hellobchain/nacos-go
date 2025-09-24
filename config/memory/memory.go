package memory

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"sync"

	"github.com/hellobchain/nacos-go/config"
)

var _ config.ConfigRepo = (*memoryRepo)(nil)

type memoryRepo struct {
	m  map[string]config.ConfigItem // key = tenant@@group@@dataId
	mu sync.RWMutex
}

// Edit implements config.ConfigRepo.
func (r *memoryRepo) Edit(ctx context.Context, dataId string, group string, tenant string, content string) error {
	key := mkKey(tenant, group, dataId)
	r.mu.Lock()
	defer r.mu.Unlock()
	item := r.m[key]
	return r.Save(ctx, config.ConfigItem{
		AppName: item.AppName,
		BetaIps: item.BetaIps,
		Content: content,
		DataId:  dataId,
		Group:   group,
		Md5:     md5Str(content),
		SrcIp:   item.SrcIp,
		SrcUser: item.SrcUser,
		Tenant:  tenant,
		Type:    item.Type,
	})
}

func New() config.ConfigRepo {
	return &memoryRepo{m: make(map[string]config.ConfigItem)}
}

func (r *memoryRepo) Save(_ context.Context, item config.ConfigItem) error {
	if item.Md5 == "" {
		item.Md5 = md5Str(item.Content)
	}
	key := mkKey(item.Tenant, item.Group, item.DataId)
	r.mu.Lock()
	r.m[key] = item
	r.mu.Unlock()
	return nil
}

func (r *memoryRepo) Get(_ context.Context, dataId, group, tenant string) (*config.ConfigItem, error) {
	key := mkKey(tenant, group, dataId)
	r.mu.RLock()
	item, ok := r.m[key]
	r.mu.RUnlock()
	if !ok {
		return nil, config.ErrNotFound
	}
	return &item, nil
}

func (r *memoryRepo) Delete(_ context.Context, dataId, group, tenant string) error {
	key := mkKey(tenant, group, dataId)
	r.mu.Lock()
	delete(r.m, key)
	r.mu.Unlock()
	return nil
}

func (r *memoryRepo) List(_ context.Context, dataId, group, tenant string) ([]config.ConfigItem, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var res []config.ConfigItem
	prefix := mkKey(tenant, group, dataId) // 支持模糊前缀
	for k, v := range r.m {
		if (tenant == "" || v.Tenant == tenant) &&
			(group == "" || v.Group == group) &&
			(dataId == "" || k == prefix || contains(v.DataId, dataId)) {
			res = append(res, v)
		}
	}
	return res, nil
}

/* ---------- 工具 ---------- */
func mkKey(tenant, group, dataId string) string {
	if tenant == "" {
		tenant = "public"
	}
	if group == "" {
		group = "DEFAULT_GROUP"
	}
	return tenant + "@@" + group + "@@" + dataId
}

func md5Str(s string) string {
	h := md5.Sum([]byte(s))
	return hex.EncodeToString(h[:])
}

func contains(a, b string) bool {
	// 简易模糊匹配
	return len(b) > 0 && len(a) >= len(b) && a[:len(b)] == b
}

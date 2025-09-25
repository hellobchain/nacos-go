package memory

import (
	"context"
	"sync"

	"github.com/hellobchain/nacos-go/tenant"
)

type memRepo struct {
	m  map[string]struct{}
	mu sync.RWMutex
}

func New() tenant.TenantRepo {
	return &memRepo{m: make(map[string]struct{})}
}

func (r *memRepo) Save(_ context.Context, name string) error {
	r.mu.Lock()
	r.m[name] = struct{}{}
	r.mu.Unlock()
	return nil
}

func (r *memRepo) Delete(_ context.Context, name string) error {
	r.mu.Lock()
	delete(r.m, name)
	r.mu.Unlock()
	return nil
}

func (r *memRepo) List(_ context.Context) ([]string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	list := make([]string, 0, len(r.m))
	for k := range r.m {
		list = append(list, k)
	}
	return list, nil
}

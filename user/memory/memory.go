package memory

import (
	"context"
	"sync"

	"github.com/hellobchain/nacos-go/user"
	"golang.org/x/crypto/bcrypt"
)

type memoryUserRepo struct {
	m  map[string]*user.User // key = username
	mu sync.RWMutex
}

// Delete implements user.UserRepo.
func (r *memoryUserRepo) Delete(ctx context.Context, username string) error {
	r.mu.Lock()
	delete(r.m, username)
	r.mu.Unlock()
	return nil
}

// List implements user.UserRepo.
func (r *memoryUserRepo) List(ctx context.Context) ([]*user.User, error) {
	r.mu.RLock()
	users := make([]*user.User, 0, len(r.m))
	for _, u := range r.m {
		users = append(users, u)
	}
	r.mu.RUnlock()
	return users, nil
}

func New() user.UserRepo {
	return &memoryUserRepo{m: make(map[string]*user.User)}
}

func (r *memoryUserRepo) Save(_ context.Context, u *user.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	// 第一次写入时把密码哈希掉
	if _, ok := r.m[u.Username]; !ok {
		hash, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		u.Password = string(hash)
	}
	r.m[u.Username] = u
	return nil
}

func (r *memoryUserRepo) GetByName(_ context.Context, username string) (*user.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if u, ok := r.m[username]; ok {
		return u, nil
	}
	return nil, user.ErrNotFound
}

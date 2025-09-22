package user

import (
	"context"
	"errors"
	"os"

	"github.com/hellobchain/nacos-go/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

var ErrNotFound = errors.New("user not found")

type AuthUserService struct {
	repo UserRepo
}

func NewAuthUserService(r UserRepo) *AuthUserService {
	return &AuthUserService{repo: r}
}

// 注册
func (s *AuthUserService) Register(ctx context.Context, username, password string) error {
	return s.repo.Save(ctx, &User{Username: username, Password: password})
}

// 登录 → 返回 JWT
func (s *AuthUserService) Login(ctx context.Context, username, password string) (string, string, error) {
	u, err := s.repo.GetByName(ctx, username)
	if err != nil {
		return "", "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return "", "", errors.New("wrong password")
	}
	uuidStr := utils.GetPureUUID()
	token, err := utils.NewSignedToken(u.ID, u.Username, "user", uuidStr, 24)
	if err != nil {
		return "", "", err
	}
	return token, uuidStr, nil
}

func InitAdminUser(svc *AuthUserService) {
	username := os.Getenv("NACOS_ADMIN")
	if username == "" {
		username = "nacos"
	}
	password := os.Getenv("NACOS_ADMIN_PASSWORD")
	if password == "" {
		password = "nacos"
	}
	// 如果存在
	_, err := svc.repo.GetByName(context.Background(), username)
	if err != nil {
		// 用户不存在 则创建
		logger.Warn("init admin user")
		if err := svc.Register(context.Background(), username, password); err != nil {
			logger.Errorf("admin user exist or err: %v", err)
		} else {
			logger.Debugf("admin user created: %s", username)
		}
	}
}

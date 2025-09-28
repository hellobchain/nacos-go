package user

import (
	"context"
	"encoding/hex"
	"errors"
	"os"

	"github.com/hellobchain/nacos-go/constant"
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
func (s *AuthUserService) Register(ctx context.Context, username, password string, role string) error {
	passwordBytes, _ := hex.DecodeString(password)
	plainBytes, err := utils.DefaultAesTool.Decrypt(passwordBytes)
	if err != nil {
		logger.Errorf("decrypt password error, err: %v", err)
		return errors.New("decrypt password error")
	}
	return s.repo.Save(ctx, &User{Username: username, Password: string(plainBytes), Role: role})
}

// 修改用户信息
func (s *AuthUserService) Update(ctx context.Context, username, password, role string) error {
	u, err := s.repo.GetByName(ctx, username)
	if err != nil {
		return err
	}
	passwordBytes, _ := hex.DecodeString(password)
	plainBytes, err := utils.DefaultAesTool.Decrypt(passwordBytes)
	if err != nil {
		logger.Errorf("decrypt password error, err: %v", err)
		return errors.New("decrypt password error")
	}
	updateUser := &User{
		ID:       u.ID,
		Username: u.Username,
		Password: string(plainBytes),
		Role:     role,
	}
	return s.repo.Save(ctx, updateUser)
}

// 获取用户列表
func (s *AuthUserService) List(ctx context.Context) ([]*User, error) {
	return s.repo.List(ctx)
}

// 删除用户
func (s *AuthUserService) Delete(ctx context.Context, username string) error {
	return s.repo.Delete(ctx, username)
}

// 登录 → 返回 JWT
func (s *AuthUserService) Login(ctx context.Context, username, password string) (string, string, error) {
	u, err := s.repo.GetByName(ctx, username)
	if err != nil {
		return "", "", err
	}
	passwordBytes, _ := hex.DecodeString(password)
	plainBytes, err := utils.DefaultAesTool.Decrypt(passwordBytes)
	if err != nil {
		logger.Errorf("decrypt password error, err: %v", err)
		return "", "", errors.New("decrypt password error")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), plainBytes); err != nil {
		logger.Errorf("wrong password, err: %v", err)
		return "", "", errors.New("wrong password")
	}
	uuidStr := utils.GetPureUUID()
	token, err := utils.NewSignedToken(u.ID, u.Username, u.Role, uuidStr, 24)
	if err != nil {
		return "", "", err
	}
	return token, uuidStr, nil
}

// 获取用户信息
func (s *AuthUserService) GetUserInfo(ctx context.Context, username string) (*User, error) {
	u, err := s.repo.GetByName(ctx, username)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (s *AuthUserService) ChangePassword(ctx context.Context, username, oldPassword, newPassword string) error {
	u, err := s.repo.GetByName(ctx, username)
	if err != nil {
		return err
	}
	oldPasswordBytes, _ := hex.DecodeString(oldPassword)
	oldPlainBytes, err := utils.DefaultAesTool.Decrypt(oldPasswordBytes)
	if err != nil {
		logger.Errorf("decrypt password error, err: %v", err)
		return errors.New("decrypt password error")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), oldPlainBytes); err != nil {
		logger.Errorf("wrong password, err: %v", err)
		return errors.New("wrong password")
	}
	newPasswordBytes, _ := hex.DecodeString(newPassword)
	newPlainBytes, err := utils.DefaultAesTool.Decrypt(newPasswordBytes)
	if err != nil {
		logger.Errorf("decrypt password error, err: %v", err)
		return errors.New("decrypt password error")
	}
	updateUser := &User{
		ID:       u.ID,
		Username: u.Username,
		Password: string(newPlainBytes),
	}
	return s.repo.Save(ctx, updateUser)
}

func InitAdminUser(svc *AuthUserService) {
	username := os.Getenv("NACOS_ADMIN")
	if username == "" {
		username = "nacos"
	}
	password := os.Getenv("NACOS_ADMIN_PASSWORD")
	if password == "" {
		password = "438c3edd02f9317d849f19d7a78ee66f"
	}
	// 如果存在
	_, err := svc.repo.GetByName(context.Background(), username)
	if err != nil {
		// 用户不存在 则创建
		logger.Warn("init admin user")
		if err := svc.Register(context.Background(), username, password, constant.ROLE_ADMIN); err != nil {
			logger.Errorf("admin user exist or err: %v", err)
		} else {
			logger.Debugf("admin user created: %s", username)
		}
	}
}

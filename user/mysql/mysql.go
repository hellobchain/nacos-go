package mysql

import (
	"context"

	"github.com/hellobchain/nacos-go/user"
	"github.com/hellobchain/wswlog/wlogging"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var logger = wlogging.MustGetFileLoggerWithoutName(nil)

type mysqlUserRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.UserRepo {
	err := db.AutoMigrate(&userPO{})
	if err != nil {
		logger.Fatal("mysql auto migrate user error:", err)
	}
	return &mysqlUserRepo{db: db}
}

type userPO struct {
	gorm.Model
	Username string `gorm:"column:username;size:64;unique;not null"`
	Password string `gorm:"column:password;size:128;not null"`
}

func (userPO) TableName() string { return "user" }

func (r *mysqlUserRepo) Save(ctx context.Context, u *user.User) error {
	po := userPO{Username: u.Username}
	// 只在首次写入时哈希
	if u.Password != "" && !isHash(u.Password) {
		hash, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		po.Password = string(hash)
	} else {
		po.Password = u.Password
	}
	return r.db.WithContext(ctx).Where("username = ?", u.Username).
		Assign(po).
		FirstOrCreate(&po).Error
}

func (r *mysqlUserRepo) GetByName(ctx context.Context, username string) (*user.User, error) {
	var po userPO
	if err := r.db.WithContext(ctx).Where("username = ?", username).First(&po).Error; err != nil {
		return nil, user.ErrNotFound
	}
	return &user.User{ID: int64(po.ID), Username: po.Username, Password: po.Password}, nil
}

func isHash(s string) bool {
	return len(s) == 60 && s[0:4] == "$2a$"
}

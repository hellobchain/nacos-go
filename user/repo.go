package user

import (
	"context"
)

type UserRepo interface {
	Save(ctx context.Context, u *User) error
	GetByName(ctx context.Context, username string) (*User, error)
	List(ctx context.Context) ([]*User, error)
	Delete(ctx context.Context, username string) error
}

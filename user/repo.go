package user

import "context"

type UserRepo interface {
	Save(ctx context.Context, u *User) error
	GetByName(ctx context.Context, username string) (*User, error)
}

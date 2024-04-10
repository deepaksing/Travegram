package store

import "context"

type Driver interface {
	Migrate(ctx context.Context) error
	CreateUser(ctx context.Context, create *User) (*User, error)
	GetUser(ctx context.Context, f *FindUser) (*User, error)
}

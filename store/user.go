package store

import (
	"context"
)

type User struct {
	ID           int32
	CreatedTs    int64
	UpdatedTs    int64
	Username     string
	PasswordHash string
	AvatarURL    string
	Country      string
	Name         string
}

type FindUser struct {
	ID       *int32
	Username *string
	Name     *string
	Country  *string
}

func (s *Store) CreateUser(ctx context.Context, userCreate *User) (*User, error) {
	user, err := s.driver.CreateUser(ctx, userCreate)
	return user, err
}

func (s *Store) GetUser(ctx context.Context, userFind *FindUser) (*User, error) {
	user, err := s.driver.GetUser(ctx, userFind)
	return user, err
}

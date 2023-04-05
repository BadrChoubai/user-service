package service

import (
	"context"
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UserRepository interface {
	AllUsers(ctx context.Context) (*[]User, error)
	SingleUserById(ctx context.Context, userId int) (*User, error)
}

type UserService interface {
	GetAllUsers(ctx context.Context) (*[]User, error)
	GetUserById(ctx context.Context, userId int) (*User, error)
}

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
	SingleUserById(ctx context.Context, userId int) (*User, error)
}

type UserService interface {
	GetUserById(ctx context.Context, userId int) (*User, error)
}

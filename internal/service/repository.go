package service

import (
	"context"

	"gorm.io/gorm"
)

const (
	QUERY_GET_USER = "SELECT * FROM users WHERE id = ?"
)

type dbRepository struct {
	database *gorm.DB
}

func NewUserRepository(dbConnection *gorm.DB) UserRepository {
	return &dbRepository{
		database: dbConnection,
	}
}

func (repository *dbRepository) SingleUserById(ctx context.Context, userId int) (*User, error) {
	user := &User{}
	repository.database.Raw(QUERY_GET_USER, userId).Scan(&user)

	return user, nil
}

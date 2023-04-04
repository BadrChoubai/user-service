package service

import (
	"context"
	"database/sql"
)

const (
	QUERY_GET_USER = "SELECT * FROM users WHERE id = ?"
)

type dbRepository struct {
	database *sql.DB
}

func NewUserRepository(dbConnection *sql.DB) UserRepository {
	return &dbRepository{
		database: dbConnection,
	}
}

func (repository *dbRepository) SingleUserById(ctx context.Context, userId int) (*User, error) {
	user := &User{}

	statement, err := repository.database.PrepareContext(ctx, QUERY_GET_USER)
	if err != nil {
		return nil, err
	}

	defer statement.Close()

	err = statement.QueryRow(ctx, userId).Scan(&user.ID, &user.Name, &user.CreatedAt, &user.UpdatedAt)

	if err != nil && err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}

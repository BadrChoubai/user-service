package infrastructure

import (
	"database/sql"
	"time"
)

func DBConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(mariadb:3306)/user_service")
	if err != nil {
		return nil, err
	}

	// Set up important parts as was told by the documentation.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	// Return our database instance.
	return db, nil
}

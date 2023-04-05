package infrastructure

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBConnection() (*gorm.DB, error) {
	dsn := "root:@tcp(mariadb:3306)/user_service"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Return our database instance.
	return db, nil
}

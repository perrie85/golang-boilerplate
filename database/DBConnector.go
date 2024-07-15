package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() gorm.DB {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	// convert to connection with dotenv
	dsn := "root:123456@tcp(db:3306)/golang_boilerplate?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	return *db
}

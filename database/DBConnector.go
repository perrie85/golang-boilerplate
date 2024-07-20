package database

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() gorm.DB {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	// convert to connection with dotenv
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	connString := user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := "root:123456@tcp(db:3306)/golang_boilerplate?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	return *db
}

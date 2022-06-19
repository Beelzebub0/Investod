package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseConnection() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		panic("Can't load .env file")
	}

	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbDatabase := os.Getenv("DB_DATABASE")

	connString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true", dbUsername, dbPassword, dbHost, dbPort, dbDatabase)

	database, err := gorm.Open(mysql.Open(connString), &gorm.Config{})
	if err != nil {
		panic("failed to connect to db")
	}

	database.AutoMigrate()

	DB = database
	return DB
}

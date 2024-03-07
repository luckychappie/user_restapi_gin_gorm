package database

import (
	"fmt"
	db_config "golang/test_rest_api/config/db"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	if db_config.DB_DRIVER == "mysql" {
		var err error
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", db_config.DB_USER, db_config.DB_PASSWORD, db_config.DB_HOST, db_config.DB_PORT, db_config.DB_NAME)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err != nil {
			panic("Connection to database fail")
		}

		DB = db
	}

	if db_config.DB_DRIVER == "postgres" {
		var err error
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", db_config.DB_HOST, db_config.DB_USER, db_config.DB_PASSWORD, db_config.DB_NAME, db_config.DB_PORT)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err != nil {
			panic("Connection to database fail")
		}

		DB = db
	}
}

package db_config

import "os"

var DB_DRIVER = "mysql"
var DB_USER = "root"
var DB_PASSWORD = "root"
var DB_NAME = "user_management"
var DB_HOST = "127.0.0.1"
var DB_PORT = "3306"

func InitialDBConfig() {
	driver_env := os.Getenv("DB_DRIVER")
	if driver_env != "" {
		DB_DRIVER = driver_env
	}

	user_env := os.Getenv("DB_USER")
	if user_env != "" {
		DB_USER = user_env
	}

	password_env := os.Getenv("DB_NAME")
	if password_env != "" {
		DB_NAME = password_env
	}

	dbname_env := os.Getenv("DB_NAME")
	if dbname_env != "" {
		DB_NAME = dbname_env
	}
}

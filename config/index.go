package config

import (
	app_config "golang/test_rest_api/config/app"
	db_config "golang/test_rest_api/config/db"
)

func InitialConfig() {
	app_config.InitialAppConig()
	db_config.InitialDBConfig()
}

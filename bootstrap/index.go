package bootstrap

import (
	"golang/test_rest_api/config"
	app_config "golang/test_rest_api/config/app"
	"golang/test_rest_api/database"
	"golang/test_rest_api/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func BootstrapApp() {
	app := gin.Default()

	err := godotenv.Load()

	if err != nil {
		log.Println("Fail to load .env file")
	}

	config.InitialConfig()

	routes.UseRouter(app)
	database.ConnectDatabase()
	app.Run(app_config.PORT)
}

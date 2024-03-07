package routes

import (
	"golang/test_rest_api/controller/file_controller"
	"golang/test_rest_api/middleware"

	"github.com/gin-gonic/gin"
)

func fileRoute(app *gin.RouterGroup) {
	r := app

	authRoute := r.Group("file", middleware.AuthMiddleware)
	authRoute.POST("/upload_file", file_controller.HandleUploadedFile)
	authRoute.DELETE("/:filename", file_controller.HandleRemoveFile)
	authRoute.POST("/middleware", middleware.UploadFile, file_controller.SendStatus)
}

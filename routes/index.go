package routes

import (
	app_config "golang/test_rest_api/config/app"
	user_controller "golang/test_rest_api/controller"
	authcontroller "golang/test_rest_api/controller/auth_controller"
	"golang/test_rest_api/middleware"

	"github.com/gin-gonic/gin"
)

func UseRouter(app *gin.Engine) {
	app.Static(app_config.STATIC_ROUTE, app_config.STATIC_DIR)

	r := app.Group("api")

	// Files
	fileRoute(r)

	// Users
	userRoute := r.Group("user")
	userRoute.GET("/", user_controller.GetAllUsers)
	userRoute.GET("/:id", user_controller.GetUserById)
	userRoute.POST("/", user_controller.CreateUser)
	userRoute.PATCH("/:id", user_controller.UpdateUser)
	userRoute.DELETE("/:id", user_controller.DeleteUser)
	userRoute.GET("/paginate", middleware.AuthMiddleware, user_controller.GetUserByPaginate)

	// Auth
	r.POST("/login", authcontroller.Login)
}

package routes

import (
	"gin-gonic-gorm/config/app_config"
	"gin-gonic-gorm/controller/book_controller"
	"gin-gonic-gorm/controller/user_controller"

	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {

	route := app.Group("api")

	route.Static(app_config.STATIC_ROUTE, app_config.STATIC_DIR)

	//user
	userRoute := route.Group("user")
	userRoute.GET("/", user_controller.GetAllUser)
	userRoute.GET("/paginate", user_controller.GetUserPaginate)
	userRoute.GET("/:id", user_controller.GetById)
	userRoute.POST("/", user_controller.Store)
	userRoute.PATCH("/:id", user_controller.Update)
	userRoute.DELETE("/:id", user_controller.Delete)


	//book
	route.GET("/book", book_controller.GetAllBook)

	v1Route(route)
}
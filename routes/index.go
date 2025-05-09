package routes

import (
	"gin-gonic-gorm/config/app_config"
	"gin-gonic-gorm/controller/book_controller"
	"gin-gonic-gorm/controller/file_controller"
	"gin-gonic-gorm/controller/user_controller"

	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {

	route := app
	route.Static(app_config.STATIC_ROUTE, app_config.STATIC_DIR)

	//user
	route.GET("/user", user_controller.GetAllUser)
	route.GET("/user/paginate", user_controller.GetUserPaginate)
	route.GET("/user/:id", user_controller.GetById)
	route.POST("/user", user_controller.Store)
	route.PATCH("/user/:id", user_controller.Update)
	route.DELETE("/user/:id", user_controller.Delete)


	//book
	route.GET("/book", book_controller.GetAllBook)

	//file
	route.POST("/file", file_controller.HandlerUploadFile)
}
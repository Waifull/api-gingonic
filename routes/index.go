package routes

import (
	"gin-gonic-gorm/controller/book_controller"
	"gin-gonic-gorm/controller/user_controller"

	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {

	route := app

	//user
	route.GET("/user", user_controller.GetAllUser)
	route.GET("/user/:id", user_controller.GetById)
	


	//book
	route.GET("/book", book_controller.GetAllBook)
}
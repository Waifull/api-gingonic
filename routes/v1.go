package routes

import (
	"gin-gonic-gorm/controller/file_controller"
	"gin-gonic-gorm/middleware"

	"github.com/gin-gonic/gin"
)

func v1Route(app *gin.RouterGroup){

	route := app

	//file
	//middleware auth route group
	authRoute := route.Group("file", middleware.AuthMiddleware)
	authRoute.POST("/", file_controller.HandlerUploadFile)
	authRoute.POST("/middleware", middleware.UploadFile, file_controller.SendStatus)
	authRoute.DELETE("/:filename", file_controller.HandleRemoveFile)
	// //contoh middleware untuk 1 route
	// //delete + middleware
	// route.DELETE("/file/:filename", middleware.AuthMiddleware, file_controller.HandleRemoveFile)
}
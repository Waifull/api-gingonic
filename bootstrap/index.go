package bootstrap

import (
	"gin-gonic-gorm/config/app_config"
	"gin-gonic-gorm/routes"

	"github.com/gin-gonic/gin"
)

func BootstrapApp() {
	app := gin.Default()

	routes.InitRoute(app)

	app.Run(app_config.PORT)
}
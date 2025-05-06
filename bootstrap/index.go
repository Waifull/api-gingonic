package bootstrap

import (
	"gin-gonic-gorm/config"
	"gin-gonic-gorm/config/app_config"
	"gin-gonic-gorm/database"
	"gin-gonic-gorm/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func BootstrapApp() {
	//Load .env FILE
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	//INIT config
	config.InitConfig()

	//DATABASE CONNECTION
	database.ConnectDatabase()

	//INIT GIN ENGINE
	app := gin.Default()

	//INIT ROUTES
	routes.InitRoute(app)

	//RUN APP
	app.Run(app_config.PORT)
}
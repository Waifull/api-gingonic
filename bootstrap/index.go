package bootstrap

import (
	"gin-gonic-gorm/config"
	"gin-gonic-gorm/config/app_config"
	"gin-gonic-gorm/config/cors_config"
	"gin-gonic-gorm/config/log_config"
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

	//Log
	log_config.DefaultLogging()

	//INIT GIN ENGINE
	app := gin.Default()

	//CORS
	app.Use(cors_config.CorsConfigContrib())
	// app.Use(cors_config.CorsConfig)

	//INIT ROUTES
	routes.InitRoute(app)

	//RUN APP
	app.Run(app_config.PORT)
}
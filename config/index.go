package config

import (
	"gin-gonic-gorm/config/app_config"
	"gin-gonic-gorm/config/db_config"
	// "gin-gonic-gorm/config/log_config"
)

func InitConfig() {
	
	app_config.InitAppConfig()
	db_config.InitDatabaseConfig()
	// log_config.DefaultLogging()
}
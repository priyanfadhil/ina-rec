package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/priyanfadhil/ina-rec/common/application"
	"github.com/priyanfadhil/ina-rec/common/config"
	"github.com/priyanfadhil/ina-rec/common/database"
	v "github.com/priyanfadhil/ina-rec/helper"
)

func main() {
	// Load environment variables
	godotenv.Load()

	// Create a Gin Gonic router
	r := gin.Default()

	// Initialize the application
	app := application.App{
		Engine:    r,
		DBManager: database.NewDatabaseManager(),
	}

	// Initialize the custom validator
	app.Engine.Use(v.CustomValidatorMiddleware(&app.Validator))

	// Initialize database with the provided configuration
	dsn := database.PostgresURI(
		config.GetConfig().DbUser,
		config.GetConfig().DbPassword,
		fmt.Sprintf(`%s:%s`, config.GetConfig().DbHost, config.GetConfig().DbPort),
		config.GetConfig().DbName)

	app.InitializeDatabase(
		dsn,
		config.GetConfig().DbMaxIdleConns,
		config.GetConfig().DbMaxOpenConns,
	)

	// Initialize the application routes and middleware
	app.Initialize()

	// Start the Gin Gonic server
	app.Start(":" + config.GetConfig().AppPort)
}

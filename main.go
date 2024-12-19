package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"goapi/src/shared/application/routes"
	"goapi/src/shared/infrastructure/di"
	"log"
)

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func main() {
	if err := initConfig(); err != nil {
		log.Fatal("error initializing config", err)
	}

	// create app
	app := fiber.New()

	// create container
	container := di.BuildContainer()

	// create routes
	routes.CreateRoutes(app, container)

	// listen
	app.Listen(":" + viper.GetString("port"))
}

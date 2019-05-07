package main

//go:generate swag init

import (
	"log"

	"github.com/Oxynger/JournalApp/db"
	"github.com/Oxynger/JournalApp/router"
	"github.com/Oxynger/JournalApp/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	swagdoc "github.com/Oxynger/JournalApp/docs"
)

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @securityDefinitions.apikey Authorization
// @in header
// @name X-Auth-Token

func init() {
	viper.AutomaticEnv()
	viper.SetDefault("mongodb_uri", "mongodb://localhost:27017")
	viper.SetDefault("port", "8080")
	viper.SetDefault("host", "localhost")

	db.Connect(viper.GetString("mongodb_uri"))

	swaggerHost := viper.GetString("host") + ":" + viper.GetString("port")
	swagdoc.SwaggerInfo.Host = swaggerHost
	swagdoc.SwaggerInfo.BasePath = "/api/v1"
	swagdoc.SwaggerInfo.Title = "API приложения для составления журналов"
	swagdoc.SwaggerInfo.Version = "1.1.1"
	swagdoc.SwaggerInfo.Description = "Это сервер предоставляющий API для сервиса электронных журналов"
}

func main() {
	app := gin.Default()
	app.Use(cors.Default())

	users := service.NewUserService()
	sessions := service.NewSessionService()
	router.V1(app.Group("/api/v1"), users, sessions)
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"log"

	"github.com/Oxynger/JournalApp/db"
	"github.com/Oxynger/JournalApp/router/v1"
	"github.com/Oxynger/JournalApp/router/v2"
	"github.com/Oxynger/JournalApp/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func init() {
	viper.AutomaticEnv()
	viper.SetDefault("mongodb_uri", "mongodb://localhost:27017")
	viper.SetDefault("port", "8080")
	viper.SetDefault("host_domain", "localhost")

	db.Connect(viper.GetString("mongodb_uri"))
}

func main() {
	app := gin.Default()
	app.Use(cors.Default())

	v1.RegisterRoutes(app.Group("/api/v1"))
	v2.RegisterRoutes(app.Group("/api/v2"), v2.Dependencies{
		service.NewUserService(),
		service.NewSessionService(),
	})

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}

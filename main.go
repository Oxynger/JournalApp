package main

import (
	"log"

	"github.com/Oxynger/JournalApp/db"
	v1 "github.com/Oxynger/JournalApp/router/v1"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func init() {
	viper.AutomaticEnv()
	viper.SetDefault("mongodb_uri", "mongodb://localhost:27017")
	viper.SetDefault("port", "8080")
	viper.SetDefault("host", "localhost")

	db.Connect(viper.GetString("mongodb_uri"))
}

func main() {
	app := gin.Default()
	app.Use(cors.Default())

	v1.RegisterRoutes(app.Group("/api/v1"))

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}

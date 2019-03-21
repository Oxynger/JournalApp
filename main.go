package main

import (
	"./config"
	"./controller"
	"./db"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	swagdoc "./docs"
)

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

var (
	conf config.Config
)

func init() {
	conf = config.New()
	db.Connect(conf.MongoURI)

	swagdoc.SwaggerInfo.Host = "localhost:" + conf.Port
	swagdoc.SwaggerInfo.BasePath = "/api/v1"
	swagdoc.SwaggerInfo.Title = "API приложения для составления журналов"
	swagdoc.SwaggerInfo.Version = "0.1.0"
	swagdoc.SwaggerInfo.Description = "Это сервер предоставляющий API для сервиса электронных журналов"

}

func main() {
	router := gin.Default()

	c := controller.New()

	v1 := router.Group("/api/v1")
	{
		itemScheme := v1.Group("/item/scheme")
		{
			itemScheme.GET("", c.TestAdd)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run()

}

package main

import (
	"log"

	"github.com/gin-contrib/cors"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"


	"github.com/Oxynger/JournalApp/api"
	"github.com/Oxynger/JournalApp/config"
	"github.com/Oxynger/JournalApp/db"
	"github.com/Oxynger/JournalApp/service"
	"github.com/gin-gonic/gin"
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
	conf := config.New()
	db.Connect(conf.MongoURI)

	swagdoc.SwaggerInfo.Host = conf.Host
	swagdoc.SwaggerInfo.BasePath = "/api/v1"
	swagdoc.SwaggerInfo.Title = "API приложения для составления журналов"
	swagdoc.SwaggerInfo.Version = "1.1.1"
	swagdoc.SwaggerInfo.Description = "Это сервер предоставляющий API для сервиса электронных журналов"
}

func main() {
	router := gin.Default()
	store := memstore.NewStore([]byte("authtest"))
	router.Use(cors.Default())
	router.Use(sessions.Sessions("X-Auth-Token", store))
	router.Use(api.Auth(service.NewUserService()))

	api.RegisterV1(router.Group("/api/v1"))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := router.Run(); err != nil {
		log.Fatal(err)
	}
}

package v2

//go:generate swag init

import (
	"github.com/Oxynger/JournalApp/api/auth"
	"github.com/Oxynger/JournalApp/api/journal"
	"github.com/Oxynger/JournalApp/api/operator"
	"github.com/Oxynger/JournalApp/api/test"
	"github.com/Oxynger/JournalApp/httputils"
	"github.com/Oxynger/JournalApp/model/user"
	"github.com/Oxynger/JournalApp/service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	swagdoc "github.com/Oxynger/JournalApp/router/v2/docs"
)

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @securityDefinitions.apikey Authorization
// @in header
// @name X-Auth-Token

// Dependencies содержит сервисы, необходимые для работы авторизации
type Dependencies struct {
	Users    *service.UserService
	Sessions *service.SessionService
}

// V2 добавляет роутинг для эндпоинтов на /api/v2
func RegisterRoutes(router *gin.RouterGroup, services Dependencies) {
	setupSwaggerDocs(router)

	router.POST("/login", auth.LogIn(services.Users, services.Sessions))
	router.POST("/user", test.AddUser(services.Users))
	router.GET("/logout", auth.LogOut(services.Sessions))

	adminGroup := router.Group("/administrator")
	{
		adminGroup.Use(auth.RequireAuthorization(services.Sessions, user.Administrator))
		adminGroup.GET("/operators", operator.ListOperators)
		adminGroup.GET("/journals", journal.ListJournals)
		adminGroup.POST("/operators", operator.AddOperator)
		adminGroup.PATCH("operators/:id", operator.UpdateOperator)
		adminGroup.POST("/reports/:id", httputils.Blank)
	}
	constGroup := router.Group("/constructor")
	{
		constGroup.Use(auth.RequireAuthorization(services.Sessions, user.Constructor))

		constGroup.GET("/schemes/journals", httputils.Blank)
		constGroup.POST("/schemes/journals", httputils.Blank)
		constGroup.PATCH("/schemes/journals", httputils.Blank)
		constGroup.DELETE("/schemes/journals/:id", httputils.Blank)
	}
	operaGroup := router.Group("/operator")
	{
		operaGroup.Use(auth.RequireAuthorization(services.Sessions, user.Operator))
		operaGroup.GET("/journals", journal.ListJournals)
		operaGroup.GET("/journals/:id", journal.ShowJournal)
		operaGroup.POST("/journals", journal.AddJournal)
		operaGroup.POST("/journals/:id/signature", journal.CloseJournal)
		operaGroup.PATCH("/journals", journal.UpdateJournal)

		operaGroup.GET("/schemes/", httputils.Blank)
	}
}

func setupSwaggerDocs(router *gin.RouterGroup) {
	swaggerHost := viper.GetString("host") + ":" + viper.GetString("port")
	swagdoc.SwaggerInfo.Host = swaggerHost
	swagdoc.SwaggerInfo.BasePath = "/api/v2"
	swagdoc.SwaggerInfo.Title = "API приложения для составления журналов"
	swagdoc.SwaggerInfo.Version = "1.1.1"
	swagdoc.SwaggerInfo.Description = "Это сервер предоставляющий API для сервиса электронных журналов"

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

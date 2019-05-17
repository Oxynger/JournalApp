package v1

import (
	"github.com/Oxynger/JournalApp/api"
	"github.com/Oxynger/JournalApp/api/itemScheme"
	"github.com/Oxynger/JournalApp/api/journal"
	"github.com/Oxynger/JournalApp/api/operator"
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

// V1 добавляет роутинг для эндпоинтов на /api/v1
func RegisterRoutes(router *gin.RouterGroup) {
	setupSwaggerDocs(router)

	itemSchemeGroup := router.Group("/scheme")
	{
		itemSchemeGroup.GET("/item", itemScheme.GetItemSchemes)
		itemSchemeGroup.GET("/item/:itemscheme_id", itemScheme.GetItemScheme)
		itemSchemeGroup.POST("/item", itemScheme.NewItemScheme)
		itemSchemeGroup.PUT("/item/:itemscheme_id", itemScheme.UpdateItemScheme)
		itemSchemeGroup.DELETE("/item/:itemscheme_id", itemScheme.DeleteItemScheme)
	}
	journalGroup := router.Group("/journal")
	{
		journalGroup.GET("", journal.ListJournals)
		journalGroup.GET(":journal_id", journal.ShowJournal)
		journalGroup.POST("", journal.AddJournal)
		journalGroup.PUT(":journal_id", journal.UpdateJournal)
		journalGroup.DELETE(":journal_id", journal.DeleteJournal)
		journalGroup.POST(":journal_id/signature", journal.CloseJournal)
	}
	operatorGroup := router.Group("/controller")
	{
		operatorGroup.GET("", operator.ListOperators)
		operatorGroup.GET(":operator_id", operator.ShowOperator)
		operatorGroup.POST("", operator.AddOperator)
		operatorGroup.PUT(":operator_id", operator.UpdateOperator)
		operatorGroup.DELETE(":operator_id", operator.DeleteOperator)
	}
	searchGroup := router.Group("/search")
	{
		searchGroup.GET("")
	}
	logs := router.Group("/logs/tabletapp")
	{
		logs.POST("", api.AddTablelog)
	}
}

func setupSwaggerDocs(router *gin.RouterGroup) {
	swaggerHost := viper.GetString("host") + ":" + viper.GetString("port")
	swagdoc.SwaggerInfo.Host = swaggerHost
	swagdoc.SwaggerInfo.BasePath = "/api/v1"
	swagdoc.SwaggerInfo.Title = "API приложения для составления журналов"
	swagdoc.SwaggerInfo.Version = "1.1.1"
	swagdoc.SwaggerInfo.Description = "Это сервер предоставляющий API для сервиса электронных журналов"

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

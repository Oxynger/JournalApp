package main

import (
	"github.com/Oxynger/JournalApp/config"
	"github.com/Oxynger/JournalApp/controller"
	"github.com/Oxynger/JournalApp/db"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	swagdoc "github.com/Oxynger/JournalApp/docs"
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

	swagdoc.SwaggerInfo.Host = conf.Host
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
		itemScheme := v1.Group("/scheme")
		{
			itemScheme.GET("/item", c.GetItemSchemes)
			itemScheme.GET("/item/:itemscheme_id", c.GetItemScheme)
			itemScheme.POST("/item", c.NewItemScheme)
			itemScheme.PUT("/item/:itemscheme_id", c.UpdateItemScheme)
			itemScheme.DELETE("/item/:itemscheme_id", c.DeleteItemScheme)
		}
		login := v1.Group("/login")
		{
			login.POST("", c.Auth)
		}
		journal := v1.Group("/journals")
		{
			journal.GET("", c.ListJouranls)
			journal.GET(":journal_id", c.ShowJournal)
			journal.GET(":journal_id/items", c.ListItemsInJournal)
			journal.POST(":journal_id/items", c.AddItemToJournal)
			journal.GET(":journal_id/items/:item_id", c.ShowItemInJournal)
			journal.PUT(":journal_id/items/:item_id", c.SaveItemInJournal)
			journal.DELETE(":journal_id/items/:item_id", c.DeleteItemFromJournal)
			journal.POST(":journal_id/items/signature", c.CloseJournal)
		}
		logs := v1.Group("/logs/tabletapp")
		{
			logs.POST("", c.AddTablelog)
		}

	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run()

}

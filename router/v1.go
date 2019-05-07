package router

import (
	"github.com/Oxynger/JournalApp/api"
	"github.com/Oxynger/JournalApp/api/auth"
	"github.com/Oxynger/JournalApp/api/itemScheme"
	"github.com/Oxynger/JournalApp/api/journal"
	"github.com/Oxynger/JournalApp/api/operator"
	"github.com/Oxynger/JournalApp/model/user"
	"github.com/Oxynger/JournalApp/service"
	"github.com/gin-gonic/gin"
)

// V1 добавляет роутинг для эндпоинтов на /api/v1
func V1(router *gin.RouterGroup, userService *service.UserService, sessionService *service.SessionService) {
	itemSchemeGroup := router.Group("/scheme")
	{
		itemSchemeGroup.Use(auth.RequireAuthorization(sessionService, user.Helpdesk))
		itemSchemeGroup.GET("/item", itemScheme.GetItemSchemes)
		itemSchemeGroup.GET("/item/:itemscheme_id", itemScheme.GetItemScheme)
		itemSchemeGroup.POST("/item", itemScheme.NewItemScheme)
		itemSchemeGroup.PUT("/item/:itemscheme_id", itemScheme.UpdateItemScheme)
		itemSchemeGroup.DELETE("/item/:itemscheme_id", itemScheme.DeleteItemScheme)
	}
	journalGroup := router.Group("/journal")
	{
		journalGroup.Use(auth.RequireAuthorization(sessionService, user.Administrator))
		journalGroup.GET("", journal.ListJournals)
		journalGroup.GET(":journal_id", journal.ShowJournal)
		journalGroup.POST("", journal.AddJournal)
		journalGroup.PUT(":journal_id", journal.UpdateJournal)
		journalGroup.DELETE(":journal_id", journal.DeleteJournal)
		journalGroup.POST(":journal_id/signature", journal.CloseJournal)
	}
	operatorGroup := router.Group("/controller")
	{
		operatorGroup.Use(auth.RequireAuthorization(sessionService, user.Administrator))
		operatorGroup.GET("", operator.ListOperators)
		operatorGroup.GET(":operator_id", operator.ShowOperator)
		operatorGroup.POST("", operator.AddOperator)
		operatorGroup.PUT(":operator_id", operator.UpdateOperator)
		operatorGroup.DELETE(":operator_id", operator.DeleteOperator)
	}
	logs := router.Group("/logs/tabletapp")
	{
		logs.POST("", api.AddTablelog)
	}
	router.POST("/login", auth.LogIn(userService, sessionService))
	router.POST("/logout", auth.LogOut(sessionService))
}

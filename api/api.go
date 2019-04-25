package api

// api — вспомогательный модуль для конструирования роутеров

import (
	"errors"
	"net/http"

	"github.com/Oxynger/JournalApp/api/itemScheme"
	"github.com/Oxynger/JournalApp/api/journal"
	"github.com/Oxynger/JournalApp/api/operator"
	"github.com/Oxynger/JournalApp/httputils"
	"github.com/gin-gonic/gin"
)

// RegisterV1 добавляет роутинг для эндпоинтов на /api/v1
func RegisterV1(v1 *gin.RouterGroup) {
	itemSchemeGroup := v1.Group("/scheme")
	{
		itemSchemeGroup.Use(requireToken)
		itemSchemeGroup.GET("/item", itemScheme.GetItemSchemes)
		itemSchemeGroup.GET("/item/:itemscheme_id", itemScheme.GetItemScheme)
		itemSchemeGroup.POST("/item", itemScheme.NewItemScheme)
		itemSchemeGroup.PUT("/item/:itemscheme_id", itemScheme.UpdateItemScheme)
		itemSchemeGroup.DELETE("/item/:itemscheme_id", itemScheme.DeleteItemScheme)
	}
	journalGroup := v1.Group("/journal")
	{
		journalGroup.Use(requireToken)
		journalGroup.GET("", journal.ListJournals)
		journalGroup.GET(":journal_id", journal.ShowJournal)
		journalGroup.POST("", journal.AddJournal)
		journalGroup.PUT(":journal_id", journal.UpdateJournal)
		journalGroup.DELETE(":journal_id", journal.DeleteJournal)
		journalGroup.POST(":journal_id/signature", journal.CloseJournal)
	}
	operatorGroup := v1.Group("/controller")
	{
		operatorGroup.Use(requireToken)
		operatorGroup.GET("", operator.ListOperators)
		operatorGroup.GET(":operator_id", operator.ShowOperator)
		operatorGroup.POST("", operator.AddOperator)
		operatorGroup.PUT(":operator_id", operator.UpdateOperator)
		operatorGroup.DELETE(":operator_id", operator.DeleteOperator)
	}
	logs := v1.Group("/logs/tabletapp")
	{
		logs.POST("", AddTablelog)
	}
}

// RegisterV2 добавляет роутинг для эндпоинтов на /api/v2
func RegisterV2(v2 *gin.RouterGroup) {

}

func requireToken(ctx *gin.Context) {
	if len(ctx.GetHeader("X-Auth-Token")) == 0 {
		httputils.NewError(ctx, http.StatusUnauthorized, errors.New("X-Auth-Token header is required"))
		ctx.Abort()
	}
	ctx.Next()
}

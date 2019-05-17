package search

import (
	"fmt"
	"net/http"

	"github.com/Oxynger/JournalApp/httputils"
	"github.com/Oxynger/JournalApp/model"
	"github.com/gin-gonic/gin"
)

// SearchJournal Поиск журнала
// @Summary Список журналов
// @Description Получение списка журналов
// @Tags Search
// @Accept  json
// @Produce  json
// @Param name body model.SearchResponse true "journal name"
// @Success 200 {object} model.JournalResponse
// @Failure 404 {object} httputils.HTTPError
// @Failure 500 {object} httputils.HTTPError
// @Security Authorization
// @Router /search/journal [post]
func SearchJournal(ctx *gin.Context) {
	fmt.Println("query")
	var query model.SearchResponse

	if err := ctx.ShouldBindJSON(&query); err != nil {
		httputils.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	journal, err := model.JournalSearch(query.ObjectName)

	if err != nil {
		httputils.NewError(ctx, http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, journal)
}

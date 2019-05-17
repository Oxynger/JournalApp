package journal

import (
	"net/http"

	"github.com/Oxynger/JournalApp/httputils"
	"github.com/Oxynger/JournalApp/model"
	"github.com/gin-gonic/gin"
)

// ListJournals Получить все журналы
// @Summary Список журналов
// @Description Получение списка журналов
// @Tags Journal
// @Accept  json
// @Produce  json
// @Success 200 {array} model.JournalResponse
// @Failure 404 {object} httputils.HTTPError
// @Failure 500 {object} httputils.HTTPError
// @Security Authorization
// @Router /journal [get]
func ListJournals(ctx *gin.Context) {

	journals, err := model.JournalsAll()

	if err != nil {
		httputils.NewError(ctx, http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, journals)
}

// ShowJournal Получение кокретного журнала
// @Summary Один журнал
// @Description Получение кокретного журнала
// @Tags Journal
// @Accept  json
// @Produce  json
// @Param journal_id path string true "Journal id"
// @Success 200 {object} model.JournalResponse
// @Failure 404 {object} httputils.HTTPError
// @Failure 500 {object} httputils.HTTPError
// @Security Authorization
// @Router /journal/{journal_id} [get]
func ShowJournal(ctx *gin.Context) {
	id := ctx.Param("journal_id")

	journal, err := model.JournalOne(id)

	if err != nil {
		httputils.NewError(ctx, http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, journal)

}

// AddJournal Добавление журнала
// @Summary Добавить журнал
// @Description Добавление журнала.
// @Tags Journal
// @Accept  json
// @Produce  json
// @Param journal body model.Journal true "journal json"
// @Success 200 {object} model.JournalResponse
// @Failure 400 {object} httputils.HTTPError
// @Failure 404 {object} httputils.HTTPError
// @Failure 500 {object} httputils.HTTPError
// @Security Authorization
// @Router /journal [post]
func AddJournal(ctx *gin.Context) {
	var journal model.Journal

	if err := ctx.ShouldBindJSON(&journal); err != nil {
		httputils.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	resaultJournal, err := model.AddJournal(journal)
	if err != nil {
		httputils.NewError(ctx, http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, resaultJournal)

}

// DeleteJournal Удаление журнала
// @Summary Удлить журнал
// @Description Удаление журнала. Установление deleted true
// @Tags Journal
// @Accept  json
// @Produce  json
// @Param journal_id path string true "Journal id"
// @Success 200 {object} model.JournalResponse
// @Failure 404 {object} httputils.HTTPError
// @Failure 500 {object} httputils.HTTPError
// @Security Authorization
// @Router /journal/{journal_id} [delete]
func DeleteJournal(ctx *gin.Context) {
	id := ctx.Param("journal_id")

	journal, err := model.JournalDelete(id)

	if err != nil {
		httputils.NewError(ctx, http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, journal)
}

// UpdateJournal Изменеие журнала
// @Summary Изменить журнал
// @Description Изменение журнала.
// @Tags Journal
// @Accept  json
// @Produce  json
// @Param journal body model.Journal true "journal json"
// @Param journal_id path string true "Journal id"
// @Success 200 {object} model.JournalResponse
// @Failure 400 {object} httputils.HTTPError
// @Failure 404 {object} httputils.HTTPError
// @Failure 500 {object} httputils.HTTPError
// @Security Authorization
// @Router /journal/{journal_id} [put]
func UpdateJournal(ctx *gin.Context) {
	id := ctx.Param("journal_id")
	var journal model.Journal

	if err := ctx.ShouldBindJSON(&journal); err != nil {
		httputils.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	resaultJournal, err := model.JournalUpdate(id, journal)

	if err != nil {
		httputils.NewError(ctx, http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, resaultJournal)

}

// CloseJournal Добавить роспись
// @Summary Добавление росписи
// @Description Добавление росписи контролера для закрытия журнала на день. Данная функция доступна только для ежедневных журналов (daily == 0). Роспись - это файл, в формате png размером 250х125.
// @Tags Journal
// @Accept  json
// @Produce  json
// @Param journal_id path string true "Journal id"
// @Success 200 {string} string "answer"
// @Failure 400 {object} httputils.HTTPError
// @Failure 404 {object} httputils.HTTPError
// @Failure 500 {object} httputils.HTTPError
// @Security Authorization
// @Router /journal/{journal_id}/signature [POST]
func CloseJournal(ctx *gin.Context) {
	httputils.Blank(ctx)
}

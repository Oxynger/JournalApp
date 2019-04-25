package api

import (
	"github.com/Oxynger/JournalApp/httputils"
	"github.com/gin-gonic/gin"
)

// AddTablelog Сохранение логов
// @Summary Сохранение логов
// @Description Сохранение логов на сервер
// @Tags Logs
// @Accept  json
// @Produce  json
// @Param errorLog query string true "Log with error"
// @Success 200 {string} string "answer"
// @Failure 400 {object} httputils.HTTPError
// @Failure 404 {object} httputils.HTTPError
// @Failure 500 {object} httputils.HTTPError
// @Router /logs/tabletapp [post]
func AddTablelog(ctx *gin.Context) {
	httputils.Blank(ctx)
}

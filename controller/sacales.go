package controller

import (
	"github.com/gin-gonic/gin"
)

// ShowScales Отдает весы с заданым id
// @Summary Показать весы
// @Description Получает id как строку
// @Tags scales
// @Accept  json
// @Produce  json
// @Param id path int true "Scales ID"
// @Success 200 {object} model.Scales
// @Router /scales/{id} [get]
func (c *Controller) ShowScales(ctx *gin.Context) {

}

// ListScales Отдает все весы
// @Summary Список акаутов
// @Description получить Весы
// @Tags scales
// @Accept  json
// @Produce  json
// @Param q query string false "name search by q"
// @Success 200 {array} model.Scales
// @Router /scales [get]
func (c *Controller) ListScales(ctx *gin.Context) {

}

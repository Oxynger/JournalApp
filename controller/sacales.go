package controller

import (
	"errors"
	"net/http"

	"../httperror"
	"../model"

	"github.com/gin-gonic/gin"
)

// ShowScales Отдает весы с заданым id
// @Summary Весы с заданым id
// @Description Принимает id как строку
// @Tags scales
// @Accept  json
// @Produce  json
// @Param id path int true "Scales ID"
// @Success 200 {object} model.Scales
// @Failure 400 {object} httperror.HTTPError
// @Failure 404 {object} httperror.HTTPError
// @Failure 500 {object} httperror.HTTPError
// @Router /scales/{id} [get]
func (c *Controller) ShowScales(ctx *gin.Context) {
	id := ctx.Param("id")

	if len(id) == 0 {
		httperror.New(ctx, http.StatusBadRequest, errors.New("Not correct id"))
		return
	}

	scales, err := model.ScalesOne(id)

	if err != nil {
		httperror.New(ctx, http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, scales)

}

// ListScales Отдает все весы
// @Summary Список весов
// @Description Получить список весов
// @Tags scales
// @Accept  json
// @Produce  json
// @Param q query string false "Поиск весов по имени из запроса q"
// @Success 200 {array} model.Scales
// @Router /scales [get]
func (c *Controller) ListScales(ctx *gin.Context) {

}

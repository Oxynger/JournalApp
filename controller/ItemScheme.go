package controller

import (
	"net/http"

	"../model"

	"github.com/gin-gonic/gin"
)

// TestAdd Вводт в базу тестовые данные
// @Summary Тест
// @Description Тест
// @Tags ItemScheme
// @Accept  json
// @Produce  json
// @Success 200 {object} model.ItemScheme
// @Failure 400 {object} httperror.HTTPError
// @Failure 404 {object} httperror.HTTPError
// @Failure 500 {object} httperror.HTTPError
// @Router /item/scheme [get]
func (c *Controller) TestAdd(ctx *gin.Context) {
	scheme := model.SomeAdd()

	ctx.JSON(http.StatusOK, scheme)
}

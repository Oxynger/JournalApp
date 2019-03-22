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
// @Failure 400 {object} httputils.HTTPError
// @Failure 404 {object} httputils.HTTPError
// @Failure 500 {object} httputils.HTTPError
// @Router /scheme/item [get]
func (c *Controller) TestAdd(ctx *gin.Context) {
	scheme := model.SomeAdd()

	ctx.JSON(http.StatusOK, scheme)
}

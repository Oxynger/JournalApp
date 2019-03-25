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

// GetItemSchemes Получить все схемы объектов
// @Summary Список схем объектов
// @Description Метод, который получает все списки объектов
// @Tags ItemScheme
// @Accept  json
// @Produce  json
// @Success 200 {object} model.ItemScheme
// @Failure 400 {object} httputils.HTTPError
// @Failure 404 {object} httputils.HTTPError
// @Failure 500 {object} httputils.HTTPError
// @Router /scheme/items [get]
func (c *Controller) GetItemSchemes(ctx *gin.Context) {

}

// GetItemScheme Получить схему объекта с id
// @Summary Схему объекта с id
// @Description Метод, который получает схему объекта с id
// @Tags ItemScheme
// @Accept  json
// @Produce  json
// @Param itemscheme_id path string true "ItemSheme id"
// @Success 200 {object} model.ItemScheme
// @Failure 400 {object} httputils.HTTPError
// @Failure 404 {object} httputils.HTTPError
// @Failure 500 {object} httputils.HTTPError
// @Router /scheme/item/{itemscheme_id} [get]
func (c *Controller) GetItemScheme(ctx *gin.Context) {

}

// NewItemScheme Создать новую схему объектов
// @Summary Новая схема объектов
// @Description Метод, который создает новую схему объектов
// @Tags ItemScheme
// @Accept  json
// @Produce  json
// @Success 200 {object} model.ItemScheme
// @Failure 400 {object} httputils.HTTPError
// @Failure 404 {object} httputils.HTTPError
// @Failure 500 {object} httputils.HTTPError
// @Router /scheme/newitem [post]
func (c *Controller) NewItemScheme(ctx *gin.Context) {

}

// UpdateItemScheme Изменить схему объектов с id
// @Summary Изменить схему объектов с id
// @Description Метод, который изменяет схему объекта с заданным id
// @Tags ItemScheme
// @Accept  json
// @Produce  json
// @Param itemscheme_id path string true "ItemSheme id"
// @Success 200 {object} model.ItemScheme
// @Failure 400 {object} httputils.HTTPError
// @Failure 404 {object} httputils.HTTPError
// @Failure 500 {object} httputils.HTTPError
// @Router /scheme/updateitem/{itemscheme_id} [put]
func (c *Controller) UpdateItemScheme(ctx *gin.Context) {

}

// DeleteItemScheme Удалить схему объектов с id
// @Summary Удалить схему объектов с id
// @Description Метод, который удаляет схему объектов с заданным id
// @Tags ItemScheme
// @Accept  json
// @Produce  json
// @Param itemscheme_id path string true "ItemSheme id"
// @Success 200 {object} model.ItemScheme
// @Failure 400 {object} httputils.HTTPError
// @Failure 404 {object} httputils.HTTPError
// @Failure 500 {object} httputils.HTTPError
// @Router /scheme/deleteitem/{itemscheme_id} [delete]
func (c *Controller) DeleteItemScheme(ctx *gin.Context) {

}

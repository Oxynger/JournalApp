package controller

import (
	"github.com/gin-gonic/gin"
)

// GetItemSchemes Получить все схемы объектов
// @Summary Список схем объектов
// @Description Метод, который получает все списки объектов
// @Tags ItemScheme
// @Accept  json
// @Produce  json
// @Success 200 {array} model.ItemScheme
// @Failure 400 {object} httputils.HTTPError
// @Failure 404 {object} httputils.HTTPError
// @Failure 500 {object} httputils.HTTPError
// @Router /scheme/item [get]
func (c *Controller) GetItemSchemes(ctx *gin.Context) {

}

// GetItemScheme Получить схему объекта с id
// @Summary Схему объекта с id
// @Description Метод, который получает схему объекта с заданным id
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
// @Param NewItemScheme body model.ItemScheme true "New Item Scheme"
// @Success 200 {object} model.ItemScheme
// @Failure 400 {object} httputils.HTTPError
// @Failure 404 {object} httputils.HTTPError
// @Failure 500 {object} httputils.HTTPError
// @Router /scheme/item [post]
func (c *Controller) NewItemScheme(ctx *gin.Context) {

}

// UpdateItemScheme Изменить схему объектов с id
// @Summary Изменить схему объектов с id
// @Description Метод, который изменяет схему объекта с заданным id
// @Tags ItemScheme
// @Accept  json
// @Produce  json
// @Param itemscheme_id path string true "ItemSheme id"
// @Param UpdateItemScheme body model.ItemScheme true "Update Item Scheme"
// @Success 200 {object} model.ItemScheme
// @Failure 400 {object} httputils.HTTPError
// @Failure 404 {object} httputils.HTTPError
// @Failure 500 {object} httputils.HTTPError
// @Router /scheme/item/{itemscheme_id} [put]
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
// @Router /scheme/item/{itemscheme_id} [delete]
func (c *Controller) DeleteItemScheme(ctx *gin.Context) {

}

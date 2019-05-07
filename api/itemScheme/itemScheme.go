package itemScheme

import (
	"net/http"

	"github.com/Oxynger/JournalApp/httputils"
	"github.com/Oxynger/JournalApp/model"
	"github.com/gin-gonic/gin"
)

// GetItemSchemes Получить все схемы объектов
// @Summary Список схем объектов
// @Description Метод, который получает все списки объектов
// @Tags ItemScheme
// @Accept  json
// @Produce  json
// @Success 200 {array} model.ItemScheme
// @Failure 404 {object} httputils.HTTPError
// @Failure 500 {object} httputils.HTTPError
// @Security Authorization
// @Router /scheme/item [get]
func GetItemSchemes(ctx *gin.Context) {
	schemes, err := model.ItemSchemeAll()
	if err != nil {
		httputils.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, schemes)
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
// @Security Authorization
// @Router /scheme/item/{itemscheme_id} [get]
func GetItemScheme(ctx *gin.Context) {
	id := ctx.Param("itemscheme_id")
	scheme, err := model.ItemSchemeOne(id)
	if err != nil {
		httputils.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, scheme)
}

// NewItemScheme Создать новую схему объектов
// @Summary Новая схема объектов
// @Description Метод, который создает новую схему объектов
// @Tags ItemScheme
// @Accept  json
// @Produce  json
// @Param NewItemScheme body model.NewItemScheme true "New Item Scheme"
// @Success 200 {object} model.NewItemScheme
// @Failure 400 {object} httputils.HTTPError
// @Failure 404 {object} httputils.HTTPError
// @Failure 500 {object} httputils.HTTPError
// @Security Authorization
// @Router /scheme/item [post]
func NewItemScheme(ctx *gin.Context) {
	var newItemScheme model.NewItemScheme
	if err := ctx.ShouldBindJSON(&newItemScheme); err != nil {
		httputils.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	if err := newItemScheme.Validation(); err != nil {
		httputils.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	err := newItemScheme.Insert()
	if err != nil {
		httputils.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, newItemScheme)
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
// @Security Authorization
// @Router /scheme/item/{itemscheme_id} [put]
func UpdateItemScheme(ctx *gin.Context) {
	id := ctx.Param("itemscheme_id")

	var updateItemScheme model.UpdateItemScheme
	if err := ctx.ShouldBindJSON(&updateItemScheme); err != nil {
		httputils.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	if err := updateItemScheme.Validation(); err != nil {
		httputils.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	err := updateItemScheme.Update(id)

	if err != nil {
		httputils.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, updateItemScheme)
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
// @Security Authorization
// @Router /scheme/item/{itemscheme_id} [delete]
func DeleteItemScheme(ctx *gin.Context) {
	id := ctx.Param("itemscheme_id")
	err := model.DeleteSchemeOne(id)
	if err != nil {
		httputils.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, id)
}

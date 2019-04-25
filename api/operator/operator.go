package operator

import (
	"net/http"

	"github.com/Oxynger/JournalApp/httputils"
	"github.com/Oxynger/JournalApp/model"
	"github.com/gin-gonic/gin"
)

// ListOperators Получить всех контроллеров
// @Summary Список контроллеров
// @Description Получение списка контроллеров
// @Tags Operator
// @Accept  json
// @Produce  json
// @Success 200 {array} model.Operator
// @Failure 404 {object} httputils.HTTPError
// @Failure 500 {object} httputils.HTTPError
// @Security Authorization
// @Router /controller [get]
func ListOperators(ctx *gin.Context) {
	operators, err := model.OperatorsAll()

	if err != nil {
		httputils.NewError(ctx, http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, operators)
}

// ShowOperator Получение конкретного контроллера
// @Summary Один контроллер
// @Description Получение конкретного контроллера
// @Tags Operator
// @Accept  json
// @Produce  json
// @Param operator_id path string true "Operator id"
// @Success 200 {object} model.Operator
// @Failure 404 {object} httputils.HTTPError
// @Failure 500 {object} httputils.HTTPError
// @Security Authorization
// @Router /controller/{operator_id} [get]
func ShowOperator(ctx *gin.Context) {
	id := ctx.Param("operator_id")

	operator, err := model.OperatorOne(id)

	if err != nil {
		httputils.NewError(ctx, http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, operator)

}

// AddOperator Добавление котнроллера
// @Summary Добавить контроллера
// @Description Добавление котнроллера.
// @Tags Operator
// @Accept  json
// @Produce  json
// @Param operator body model.Operator true "operator json"
// @Success 200 {object} model.Operator
// @Failure 400 {object} httputils.HTTPError
// @Failure 404 {object} httputils.HTTPError
// @Failure 500 {object} httputils.HTTPError
// @Security Authorization
// @Router /controller [post]
func AddOperator(ctx *gin.Context) {
	var operator model.Operator

	if err := ctx.ShouldBindJSON(&operator); err != nil {
		httputils.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	err := operator.HashPassword()

	if err != nil {
		httputils.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	resaultOperator, err := model.AddOperator(operator)

	if err != nil {
		httputils.NewError(ctx, http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, resaultOperator)

}

// DeleteOperator Удаление контроллера
// @Summary Удлить контроллер
// @Description Удаление контроллера. Установление deleted true
// @Tags Operator
// @Accept  json
// @Produce  json
// @Param operator_id path string true "Operator id"
// @Success 200 {object} model.Operator
// @Failure 404 {object} httputils.HTTPError
// @Failure 500 {object} httputils.HTTPError
// @Security Authorization
// @Router /controller/{operator_id} [delete]
func DeleteOperator(ctx *gin.Context) {
	id := ctx.Param("operator_id")

	operator, err := model.OperatorDelete(id)

	if err != nil {
		httputils.NewError(ctx, http.StatusNotFound, err)
	}

	ctx.JSON(http.StatusOK, operator)
}

// UpdateOperator Изменеие котнроллера
// @Summary Изменить котнроллер
// @Description Изменение котнроллера.
// @Tags Operator
// @Accept  json
// @Produce  json
// @Param operator body model.Operator true "operator json"
// @Param operator_id path string true "Operator id"
// @Success 200 {object} model.Operator
// @Failure 400 {object} httputils.HTTPError
// @Failure 404 {object} httputils.HTTPError
// @Failure 500 {object} httputils.HTTPError
// @Security Authorization
// @Router /controller/{operator_id} [put]
func UpdateOperator(ctx *gin.Context) {
	id := ctx.Param("operator_id")
	var operator model.Operator

	if err := ctx.ShouldBindJSON(&operator); err != nil {
		httputils.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	err := operator.HashPassword()

	if err != nil {
		httputils.NewError(ctx, http.StatusOK, err)
		return
	}

	resaultOperator, err := model.OperatorUpdate(id, operator)

	if err != nil {
		httputils.NewError(ctx, http.StatusNotFound, err)
	}

	ctx.JSON(http.StatusOK, resaultOperator)
}

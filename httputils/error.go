package httputils

import (
	"github.com/gin-gonic/gin"
)

// New Конструктор ошибки
func NewError(ctx *gin.Context, status int, err error) {
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}

	ctx.JSON(status, er)
}

// HTTPError Объект ошибки
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

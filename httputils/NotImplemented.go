package httputils

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

const message string = "Этот метод API ещё не реализован"

// Blank создание заглушки
func Blank(ctx *gin.Context) {
	NewError(ctx, http.StatusOK, errors.New(message))
}

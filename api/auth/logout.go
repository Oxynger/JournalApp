package auth

import (
	"net/http"

	"github.com/Oxynger/JournalApp/service"
	"github.com/gin-gonic/gin"
)

// LogOut используется для завершения сессии
// @Summary Регистрация на сервере (WIP)
// @Description Завершение сессии пользователей
// @Accept json
// @Produce json
// @Param user body user.User true "user json"
// @Success 200 {string} object
// @Failure 400 {object} httputils.HTTPError
// @Failure 500 {object} httputils.HTTPError
// @Router /logout [get]
func LogOut(srv *service.SessionService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("X-Auth-Token")
		srv.InvalidateToken(token)
		ctx.Status(http.StatusOK)
	}
}

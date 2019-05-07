package auth

import (
	"errors"
	"net/http"

	"github.com/Oxynger/JournalApp/httputils"
	"github.com/Oxynger/JournalApp/model/user"
	"github.com/Oxynger/JournalApp/service"
	"github.com/gin-gonic/gin"
)

func RequireAuthorization(srv *service.SessionService, requiredRole user.Role) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("X-Auth-Token")
		if len(token) == 0 {
			httputils.NewError(ctx, http.StatusUnauthorized, errors.New("X-Auth-Token header is required"))
			ctx.Abort()
		}
		if srv.VerifyToken(token) {
			httputils.NewError(ctx, http.StatusUnauthorized, errors.New("Token is expired or invalid"))
			ctx.Abort()
		}
		ctx.Next()
	}
}

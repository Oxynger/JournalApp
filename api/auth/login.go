package auth

import (
	"net/http"

	"github.com/Oxynger/JournalApp/httputils"
	"github.com/Oxynger/JournalApp/model/user"
	"github.com/Oxynger/JournalApp/service"
	"github.com/gin-gonic/gin"
)

type Token struct {
	Token    string `json:"token"`
	ExpireAt int64  `json:"expiresAt"`
}

// LogIn используется для авторизации
// @Summary Авторизация на сервере (WIP)
// @Description Авторизация на сервере
// @Accept json
// @Produce json
// @Param credentials body user.Credentials true "credentials json"
// @Success 200 {string} Token
// @Failure 401 {object} httputils.HTTPError
// @Failure 500 {object} httputils.HTTPError
// @Router /login [post]
func LogIn(users *service.UserService, sessions *service.SessionService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var creds user.Credentials
		if err := ctx.ShouldBind(&creds); err != nil {
			httputils.NewError(ctx, http.StatusBadRequest, err)
			return
		}

		usr, err := users.Authenticate(creds)
		if err != nil {
			httputils.NewError(ctx, http.StatusUnauthorized, err)
		}

		session, err := sessions.CreateSession(usr)
		if err != nil {
			httputils.NewError(ctx, http.StatusUnauthorized, err)
		}

		ctx.JSON(http.StatusOK, Token{
			Token:    session.Token,
			ExpireAt: session.ExpireAt,
		})
	}
}

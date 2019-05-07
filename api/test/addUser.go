package test

import (
	"net/http"

	"github.com/Oxynger/JournalApp/httputils"
	"github.com/Oxynger/JournalApp/model/user"
	"github.com/Oxynger/JournalApp/service"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// AddUser добавляет пользователей без авторизации
// и предназначена только для тестирования
// @Summary Добавляет пользователей без авторизации
// @Description Авторизация на сервере (WIP)
// @Accept json
// @Produce json
// @Param credentials body user.Credentials true "credentials json"
// @Success 200 {string} user.User
// @Failure 400 {object} httputils.HTTPError
// @Failure 500 {object} httputils.HTTPError
// @Router /test/user [post]
func AddUser(users *service.UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var creds user.Credentials
		if err := ctx.ShouldBind(&creds); err != nil {
			httputils.NewError(ctx, http.StatusBadRequest, err)
			return
		}

		usr := user.User{
			ID:       uuid.NewV4(),
			Username: creds.Username,
			Password: creds.Password,
			Role:     user.Administrator,
		}

		err := users.Create(usr)
		if err != nil {
			httputils.NewError(ctx, http.StatusUnauthorized, err)
		}

		ctx.JSON(http.StatusOK, usr)
	}
}

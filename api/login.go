package api

import (
	"net/http"

	"github.com/Oxynger/JournalApp/httputils"
	"github.com/Oxynger/JournalApp/model/user"
	"github.com/gin-gonic/gin"
)

// Login используется для авторизации
// @Summary Авторизация на сервере (WIP)
// @Description Авторизация на сервере
// @Accept json
// @Produce json
// @Param credentials body user.Credentials true "credentials json"
// @Success 200 {string} object
// @Failure 400 {object} httputils.HTTPError
// @Failure 500 {object} httputils.HTTPError
// @Router /login [post]
func Login(ctx *gin.Context) {
	var creds user.Credentials
	if err := ctx.ShouldBind(&creds); err != nil {
		httputils.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": ""})
}

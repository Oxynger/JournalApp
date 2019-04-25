package api

import (
	"net/http"

	"github.com/Oxynger/JournalApp/db"
	"github.com/Oxynger/JournalApp/httputils"
	"github.com/Oxynger/JournalApp/model/user"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// Login используется для авторизации
// @Summary Авторизация на сервере
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

	//users := userCollection()

	ctx.JSON(http.StatusOK, gin.H{"token": ""})
}

func userCollection() *mongo.Collection {
	client := db.Client()
	return client.Database("test").Collection("Users")
}

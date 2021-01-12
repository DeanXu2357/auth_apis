package handlers_v1

import (
	"auth/internal/helpers"
	"auth/internal/models"
	"github.com/gin-gonic/gin"
)

func ShowUser(c *gin.Context) {
	auth := helpers.GetAuth(c)
	db := helpers.GetDB(c)

	user := models.User{ID: auth.UserID}
	if err := db.First(&user).Error; err != nil {
		helpers.GenerateResponse(c, helpers.ReturnBadRequest, map[string]string{"detail": err.Error()})
		return
	}

	helpers.GenerateResponse(c, helpers.ReturnOK, user)
	return
}

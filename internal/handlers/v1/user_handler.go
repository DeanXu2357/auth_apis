package handlers_v1

import (
	"auth/internal/helpers"
	"auth/internal/models"
	"github.com/gin-gonic/gin"
)

// ShowUser godoc
// @Summary User profile
// @Description Get user profile
// @Tags User
// @Accept  json
// @Produce  json
// @Success 200 {object} helpers.ResponseContent{items=models.User} string "{"status":200, "msg":ok}"
// @Failure 400 {object} helpers.ResponseContent string "40000: invalid token"
// @Failure 500
// @Security ApiKeyAuth
// @Param Authorization header string true "With the bearer started"
// @Router /api/v1/user [get]
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

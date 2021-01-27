package handlers_v1

import (
	"auth/internal/config"
	"auth/internal/helpers"
	"auth/internal/models"
	"auth/internal/services"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"time"
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

func RefreshToken(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	tokenString := authHeader[len("Bearer "):]
	db := helpers.GetDB(c)

	authToken, err := services.DecodeLoginToken(tokenString, db.Session(&gorm.Session{NewDB: true}))
	if err != nil {
		if !errors.Is(err, services.ErrorTokenExpired) {
			helpers.GenerateResponse(c, helpers.ReturnValidationFailed, map[string]string{"detail": err.Error()})
			return
		}
	}

	if authToken.IsRevoked() {
		helpers.GenerateResponse(c, helpers.ReturnValidationFailed, map[string]string{"detail": "token_revoked"})
		return
	}

	// check if out of refresh limit
	refreshExpire := authToken.CreatedAt.Add(time.Duration(config.LoginAuth.RefreshExpire) * time.Second)
	if time.Now().After(refreshExpire) {
		helpers.GenerateResponse(c, helpers.ReturnValidationFailed, map[string]string{"detail": "out_of_refresh_time"})
		return
	}

	user := models.User{ID: authToken.UserID}
	if err := db.First(&user).Error; err != nil {
		helpers.GenerateResponse(c, helpers.ReturnNotExist, map[string]string{"detail": "user not exist"})
		return
	}

	tx := db.Session(&gorm.Session{SkipDefaultTransaction: true, NewDB: true})
	defer func() {
		if r := recover(); r != nil {
			log.Print(r.(error))
			tx.Rollback()
		}
	}()

	// generate new token
	tokenString, err = services.GenerateLoginToken(user, tx, "refresh_token")
	if err != nil {
		tx.Rollback()
		helpers.GenerateResponse(c, helpers.ReturnInternalError, map[string]string{"detail": err.Error()})
		return
	}

	// revoke old token
	if err = authToken.DoRevoked(tx); err != nil {
		tx.Rollback()
		helpers.GenerateResponse(c, helpers.ReturnInternalError, map[string]string{"detail": err.Error()})
		return
	}

	tx.Commit()

	helpers.GenerateResponse(c, helpers.ReturnOK, map[string]string{"token": tokenString})
	return
}

package handlers_v1

import (
	"auth/internal/helpers"
	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	helpers.GenerateResponse(c, helpers.ReturnOK, nil)
}

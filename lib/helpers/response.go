package helpers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Response struct {
	status int
	msg    string
	http   int
}

var (
	ReturnOK               = Response{status: 200, msg: "ok", http: http.StatusOK}
	ReturnInvalidToken     = Response{status: 40101, msg: "Invalid token", http: http.StatusUnauthorized}
	ReturnTokenExpire      = Response{status: 40102, msg: "Expired token", http: http.StatusUnauthorized}
	ReturnValidationFailed = Response{status: 40022, msg: "Validation failed", http: http.StatusBadRequest}
	ReturnResourceNotFound = Response{status: 404, msg: "Resource not found", http: http.StatusNotFound}
	ReturnNotExist         = Response{status: 40004, msg: "Target not exist", http: http.StatusBadRequest}
	ReturnDuplicate        = Response{status: 40009, msg: "Already exists", http: http.StatusBadRequest}
	ReturnInternalError    = Response{status: 500, msg: "Internal error", http: http.StatusInternalServerError}
)

func GenerateResponse(c *gin.Context, r Response, items interface{}) {
	c.JSON(r.http, gin.H{"status": r.status, "msg": r.msg, "items": items})
}

func GeneratePagination(c *gin.Context, items interface{}, total int) {
	per, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	current, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	//from := (current-1)*per + 1
	//to := from +

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "OK",
		"items":  items,
		//"from":         from,
		//"to":           "",
		"total":        total,
		"per_page":     per,
		"current_page": current,
		//"last_page":    "",
	})
}

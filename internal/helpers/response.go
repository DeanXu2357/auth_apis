package helpers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Response struct {
	Http    int
	Content ResponseContent
}

// ResponseContent example
type ResponseContent struct {
	Status int         `json:"status" example:"200"`
	Msg    string      `json:"msg" example:"ok"`
	Items  interface{} `json:"items"`
}

var (
	ReturnOK               = Response{Http: http.StatusOK, Content: ResponseContent{Status: 200, Msg: "ok"}}
	ReturnBadRequest       = Response{Http: http.StatusBadRequest, Content: ResponseContent{Status: 40000, Msg: "bad request"}}
	ReturnInvalidToken     = Response{Http: http.StatusUnauthorized, Content: ResponseContent{Status: 40101, Msg: "Invalid token"}}
	ReturnTokenExpire      = Response{Http: http.StatusUnauthorized, Content: ResponseContent{Status: 40102, Msg: "Expired token"}}
	ReturnValidationFailed = Response{Http: http.StatusBadRequest, Content: ResponseContent{Status: 40022, Msg: "Validation failed"}}
	ReturnResourceNotFound = Response{Http: http.StatusNotFound, Content: ResponseContent{Status: 404, Msg: "Resource not found"}}
	ReturnNotExist         = Response{Http: http.StatusBadRequest, Content: ResponseContent{Status: 40004, Msg: "Target not exist"}}
	ReturnDuplicate        = Response{Http: http.StatusBadRequest, Content: ResponseContent{Status: 40009, Msg: "Already exists"}}
	ReturnInternalError    = Response{Http: http.StatusInternalServerError, Content: ResponseContent{Status: 500, Msg: "Internal error"}}
)

func GenerateResponse(c *gin.Context, r Response, items interface{}) {
	r.Content.Items = items
	c.JSON(r.Http, r.Content)
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

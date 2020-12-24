package helpers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HttpResponse struct {
	ctx *gin.Context
}

func NewHttpResponse(c *gin.Context) *HttpResponse {
	return &HttpResponse{ctx: c}
}

func (r *HttpResponse) Ok(items interface{}) {
	dataString, err := json.Marshal(items)
	if err != nil {
		// todo :
	}

	r.ctx.JSON(http.StatusOK, gin.H{"status": 20000, "msg": "OK", "items": dataString})
}

func (r *HttpResponse) Paginate(items interface{}) {
	r.ctx.JSON(http.StatusOK, gin.H{
		"status":       20000,
		"msg":          "OK",
		"items":        items,
		"from":         "",
		"to":           "",
		"total":        "",
		"per_page":     "",
		"current_page": "",
		"last_page":    "",
	})
}

func (r *HttpResponse) Error(e Error) {
}

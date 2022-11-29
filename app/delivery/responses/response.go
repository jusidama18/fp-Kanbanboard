package responses

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string      `json:"message"`
	Error   interface{} `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func BadRequestError(c *gin.Context, err interface{}) {
	resp := &Response{
		Message: "BAD_REQUEST",
		Error:   err,
	}

	c.AbortWithStatusJSON(http.StatusBadRequest, resp)

}

func InternalServerError(c *gin.Context, err interface{}) {
	resp := &Response{
		Message: "INTERNAL_SERVER_ERROR",
		Error:   err,
	}

	c.AbortWithStatusJSON(http.StatusBadRequest, resp)
}

func UnauthorizedRequest(c *gin.Context, msg string) {
	resp := &Response{
		Message: fmt.Sprintf("UNAUTHORIZED_REQUEST: %s", msg),
	}
	c.AbortWithStatusJSON(http.StatusUnauthorized, resp)
}

func Success(c *gin.Context, statusCode int, msg string, data ...interface{}) {
	resp := &Response{
		Data:    data,
		Message: msg,
	}
	c.JSON(statusCode, resp)
}

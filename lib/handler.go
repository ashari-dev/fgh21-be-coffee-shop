package lib

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandlerOK(c *gin.Context, msg string, data any, pageInfo any) {
	c.JSON(http.StatusOK, Respont{
		Success:  true,
		Message:  msg,
		PageInfo: pageInfo,
		Result:   data,
	})
}

func HandlerNotfound(c *gin.Context, msg string) {
	c.JSON(http.StatusNotFound, Respont{
		Success:  false,
		Message:  msg,
	})
}

func HandlerUnauthorized(c *gin.Context, msg string) {
	c.JSON(http.StatusUnauthorized, Respont{
		Success:  false,
		Message:  msg,
	})
}

func HandlerBadReq(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest, Respont{
		Success:  false,
		Message:  msg,
	})
}

func HandlerMaxFile(c *gin.Context, msg string){
	c.JSON(http.StatusRequestEntityTooLarge, Respont{
		Success: false,
		Message: msg,
	})
}
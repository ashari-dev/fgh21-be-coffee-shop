package middlewares

import (
	"RGT/konis/lib"

	"github.com/gin-gonic/gin"
)

func tokenFailed(c *gin.Context) {
	if e := recover(); e != nil {
		lib.HandlerUnauthorized(c, "Unauthorized")
		c.Abort()
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer tokenFailed(ctx)
		token := ctx.GetHeader("Authorization")[7:]
		isValidated, userId := lib.ValidateToken(token)

		if isValidated {
			ctx.Set("UserId", userId)
			ctx.Next()
		} else {
			panic("Err: token Invalid")
		}
	}
}

package routers

import (
	"RGT/konis/controllers"
	"RGT/konis/middlewares"

	"github.com/gin-gonic/gin"
)

func CartsRouters(rg *gin.RouterGroup) {
	rg.GET("", middlewares.AuthMiddleware(), controllers.ListAllCarts)
	rg.DELETE("", middlewares.AuthMiddleware(), controllers.DeleteCartsByUserId)
	rg.POST("/:id", middlewares.AuthMiddleware(), controllers.CreateOneCarts)
}

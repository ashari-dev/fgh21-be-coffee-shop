package routers

import (
	"RGT/konis/controllers"
	"RGT/konis/middlewares"

	"github.com/gin-gonic/gin"
)

func CartsRouters(rg *gin.RouterGroup) {
	rg.GET("", middlewares.AuthMiddleware(), controllers.ListAllCarts)
	rg.DELETE("/:id", controllers.DeleteOneCarts)
	rg.POST("", middlewares.AuthMiddleware(), controllers.CreateOneCarts)
}

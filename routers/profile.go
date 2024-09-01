package routers

import (
	"RGT/konis/controllers"
	"RGT/konis/middlewares"

	"github.com/gin-gonic/gin"
)

func ProfileRouters(rg *gin.RouterGroup) {
	rg.GET("/",middlewares.AuthMiddleware(), controllers.FindProfileById)
	rg.PATCH("/", middlewares.AuthMiddleware(), controllers.UpdateProfile)
	rg.DELETE("/:id", controllers.DeleteProfile)
}

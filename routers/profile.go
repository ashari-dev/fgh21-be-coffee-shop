package routers

import (
	"RGT/konis/controllers"
	"RGT/konis/middlewares"

	"github.com/gin-gonic/gin"
)

func ProfileRouters(rg *gin.RouterGroup) {
	rg.GET("", controllers.GetALLProfiles)
	rg.PATCH("/:id", controllers.UpdateProfile)
	rg.GET("/", middlewares.AuthMiddleware(), controllers.FindProfileById)
	rg.DELETE("/:id", controllers.DeleteProfile)
	rg.PATCH("/img", middlewares.AuthMiddleware(), controllers.UploadProfileImage)
}

package routers

import (
	"RGT/konis/controllers"
	"RGT/konis/middlewares"

	"github.com/gin-gonic/gin"
)

func ProfileRouters(rg *gin.RouterGroup) {
	rg.Use(middlewares.AuthMiddleware())
	// rg.GET("/",middlewares.AuthMiddleware(), controllers.FindProfileById)
	rg.PATCH("/", controllers.UpdateProfile)
	// rg.GET("", controllers.GetALLProfiles)
	rg.GET("", controllers.FindProfileById)
	rg.POST("", controllers.CreateProfileJoinUser)
	rg.PATCH("/:id", controllers.UpdateProfile)
	rg.GET("/:id", controllers.FindProfileById)
	rg.DELETE("/:id", controllers.DeleteProfile)
	rg.PATCH("/img", controllers.UploadProfileImage)
}

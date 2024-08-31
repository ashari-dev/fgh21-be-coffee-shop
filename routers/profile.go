package routers

import (
	"RGT/konis/controllers"

	"github.com/gin-gonic/gin"
)

func ProfileRouters(rg *gin.RouterGroup) {
	rg.PATCH("/:id", controllers.UpdateProfile)
	rg.GET("/:id", controllers.FindProfileById)
	rg.DELETE("/:id", controllers.DeleteProfile)
}

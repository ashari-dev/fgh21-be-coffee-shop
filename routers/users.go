package routers

import (
	"RGT/konis/controllers"

	"github.com/gin-gonic/gin"
)

func UserRouters(rg *gin.RouterGroup) {
	rg.GET("", controllers.GetAllUsers)
}

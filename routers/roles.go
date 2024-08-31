package routers

import (
	"RGT/konis/controllers"

	"github.com/gin-gonic/gin"
)

func RolesRouters(rg *gin.RouterGroup) {
	rg.GET("", controllers.GetAllRoles)
	rg.GET("/:id", controllers.GetOneRoles)

}

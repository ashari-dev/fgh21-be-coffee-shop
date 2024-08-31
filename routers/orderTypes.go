package routers

import (
	"RGT/konis/controllers"

	"github.com/gin-gonic/gin"
)

func OrderTypesRouters(rg *gin.RouterGroup) {
	rg.GET("", controllers.SeeAllOrderTypes)
	rg.GET("/:id", controllers.SeeOrderTypesById)
}

package routers

import (
	"RGT/konis/controllers"

	"github.com/gin-gonic/gin"
)
func ProductRouters(rg *gin.RouterGroup) {
	rg.GET("", controllers.GetAllProduct)
	rg.GET("/:id", controllers.GEtOneProduct)
}
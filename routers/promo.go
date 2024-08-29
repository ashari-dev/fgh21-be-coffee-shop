package routers

import (
	"RGT/konis/controllers"

	"github.com/gin-gonic/gin"
)

func PromoRouters(rg *gin.RouterGroup) {
	rg.GET("", controllers.ListAllPromo)
	rg.GET("/:id",controllers.ListPromoById)
}
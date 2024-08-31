package routers

import (
	"RGT/konis/controllers"

	"github.com/gin-gonic/gin"
)

func ProductOrderTypeRouters (r *gin.RouterGroup) {
	r.GET("", controllers.ListAllProductOrderType)
	r.POST("", controllers.CreateNewProductOrderType)
	r.GET("/:id", controllers.DetailProductOrderTypeById)
	r.PATCH("/:id", controllers.UpdateProductOrderType)
	r.DELETE("/:id", controllers.DeleteProductOrderType)
}
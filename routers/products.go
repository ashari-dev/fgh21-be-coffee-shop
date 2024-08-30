package routers

import (
	"RGT/konis/controllers"

	"github.com/gin-gonic/gin"
)

func ProductsRouters(r *gin.RouterGroup) {
	r.GET("", controllers.ListAllProducts)
	r.POST("", controllers.CreateProduct)
	r.GET("/:id", controllers.ListProductById)
	r.PATCH("/:id", controllers.UpdateProduct)
	r.DELETE("/:id", controllers.DeleteProduct)
	r.GET("/productSizes", controllers.ListAllProductsSize)
	r.GET("/productSizes/:id", controllers.GetProductsById)

}

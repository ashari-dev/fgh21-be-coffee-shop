package routers

import (
	"RGT/konis/controllers"

	"github.com/gin-gonic/gin"
)

func ProductsRouters(r *gin.RouterGroup) {
	r.GET("", controllers.ListAllProducts)
	r.GET("/", controllers.ListProductsWithPagination)
	r.POST("", controllers.CreateProduct)
	r.GET("/:id", controllers.ListProductById)
	r.PATCH("/:id", controllers.UpdateProduct)
	r.DELETE("/:id", controllers.DeleteProduct)
	r.GET("/productSizes", controllers.ListAllProductsSize)
	r.GET("/productSizes/:id", controllers.GetProductsById)
	r.GET("/variant", controllers.ListAllProductVariant)
	r.GET("/variant/:id", controllers.ListProductVariantById)
	r.GET("/categoryproducts/", controllers.ListAllCategoryProduct)
	r.GET("/categoryproducts/:id", controllers.GetCategoryProductByCategoryId)

}

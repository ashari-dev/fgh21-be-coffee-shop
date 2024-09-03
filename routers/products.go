package routers

import (
	"RGT/konis/controllers"

	"github.com/gin-gonic/gin"
)

func ProductsRouters(r *gin.RouterGroup) {

	r.GET("", controllers.ListAllProducts)
	r.POST("", controllers.CreateProduct)
	r.GET("/", controllers.ListProductsWithPagination)
	r.GET("/filter/", controllers.ListAllFilterProductsWithPagination)
	r.GET("/our-product/", controllers.ListAllOurProductsWithPagination)
	r.POST("", middlewares.AuthMiddleware(), controllers.CreateProduct)
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

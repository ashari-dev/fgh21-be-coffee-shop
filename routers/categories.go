package routers

import (
	"RGT/konis/controllers"

	"github.com/gin-gonic/gin"
)

func CategoriesRouters(rg *gin.RouterGroup) {
	rg.GET("", controllers.GetALLCategories)
}

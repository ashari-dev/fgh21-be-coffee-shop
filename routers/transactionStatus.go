package routers

import (
	"RGT/konis/controllers"

	"github.com/gin-gonic/gin"
)

func TransactionStatusRouters(rg *gin.RouterGroup) {
	rg.GET("/:id", controllers.GetTRansactionStatusById)
	rg.GET("", controllers.GetALLTransactionStatus)
}

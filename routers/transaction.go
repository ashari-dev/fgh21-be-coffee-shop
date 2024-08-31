package routers

import (
	"RGT/konis/controllers"

	"github.com/gin-gonic/gin"
)

func TransactionRouters(rg *gin.RouterGroup) {
	rg.GET("/:id", controllers.GetTransactionDetailById)
}

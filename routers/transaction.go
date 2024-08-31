package routers

import (
	"RGT/konis/controllers"


	"RGT/konis/middlewares"


	"github.com/gin-gonic/gin"
)


func TransactionRouters(rg *gin.RouterGroup) {
	rg.POST("",middlewares.AuthMiddleware(), controllers.CreateTransaction)
	rg.GET("/:id", controllers.GetTransactionDetailById)
}


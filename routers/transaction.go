package routers

import (
	"RGT/konis/controllers"

	"RGT/konis/middlewares"

	"github.com/gin-gonic/gin"
)

func TransactionRouters(rg *gin.RouterGroup) {
	rg.POST("", middlewares.AuthMiddleware(), controllers.CreateTransaction)
	rg.GET("", middlewares.AuthMiddleware(), controllers.GetAllTransactionByUserId)
	rg.POST("/:id", controllers.CreateTransactionDetails)
	rg.GET("/:id", controllers.GetTransactionDetailById)
	rg.GET("/products/:id", controllers.GetTransactionProductById)
}

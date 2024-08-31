package routers

import (
	"RGT/konis/controllers"

	"github.com/gin-gonic/gin"
)

func UserRouters(rg *gin.RouterGroup) {
	rg.GET("", controllers.GetAllUsers)
	rg.POST("", controllers.CreateUser)
	rg.GET("/:id", controllers.GetUserById)
	rg.PATCH("/:id", controllers.UpdateUserById)
	rg.DELETE("/:id", controllers.DeleteUserById)
}

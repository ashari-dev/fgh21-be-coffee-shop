package routers

import (
	"RGT/konis/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRouters(rg *gin.RouterGroup) {

	rg.POST("/login", controllers.AuthLogin)
	rg.POST("/register", controllers.RegisterLogin)

}

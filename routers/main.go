package routers

import "github.com/gin-gonic/gin"

func RouterCombain(r *gin.Engine) {
	UserRouters(r.Group("/user"))
	ProductsRouters(r.Group("/products"))
}

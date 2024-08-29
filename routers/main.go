package routers

import "github.com/gin-gonic/gin"

func RouterCombain(r *gin.Engine) {
	UserRouters(r.Group("/user"))
	RolesRouters(r.Group("/roles"))
}

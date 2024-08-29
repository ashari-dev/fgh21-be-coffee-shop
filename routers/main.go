package routers

import "github.com/gin-gonic/gin"

func RouterCombine(r *gin.Engine) {
	UserRouters(r.Group("/user"))
	AuthRouters(r.Group("/auth"))
	ProfileRouters(r.Group("/profile"))
}

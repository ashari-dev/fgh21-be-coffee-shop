package routers

import "github.com/gin-gonic/gin"


func RouterCombain(r *gin.Engine) {
	RolesRouters(r.Group("/roles"))
	UserRouters(r.Group("/users"))

	CategoriesRouters(r.Group("categories"))

	AuthRouters(r.Group("/auth"))
  ProfileRouters(r.Group("/profile"))


}

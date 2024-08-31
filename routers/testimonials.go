package routers

import (
	"RGT/konis/controllers"

	"github.com/gin-gonic/gin"
)

func TestimonialsRouters(rg *gin.RouterGroup) {
	rg.GET("", controllers.GetAllTestimonials)
}

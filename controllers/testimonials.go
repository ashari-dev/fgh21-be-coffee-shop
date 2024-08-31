package controllers

import (
	"RGT/konis/lib"
	"RGT/konis/repository"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllTestimonials(c *gin.Context) {
	results, err := repository.FindAllTestimonials()
	fmt.Println(err)
	if err != nil {
		lib.HandlerNotfound(c, "Testimonials not found")
		return
	}
	c.JSON(http.StatusOK, lib.Respont{
		Success: true,
		Message: "List All Testimonials",
		Result:  results,
	})
}

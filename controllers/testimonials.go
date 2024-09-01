package controllers

import (
	"RGT/konis/lib"
	"RGT/konis/repository"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllTestimonials(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	if page%4 == 0 {
		page = 4
	}
	if page > 4 {
		page = page % 4
	}
	if page < 1 {
		page = 1
	}
	results, err := repository.FindAllTestimonials(page)
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

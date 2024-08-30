package controllers

import (
	"RGT/konis/lib"
	"RGT/konis/repository"

	"github.com/gin-gonic/gin"
)

func GetALLCategories(c *gin.Context) {
	products := repository.FindAllCategories()

	lib.HandlerOK(c, "List All Category", products, nil)
}

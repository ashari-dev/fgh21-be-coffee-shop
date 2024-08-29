package controllers

import (
	"RGT/konis/lib"
	"RGT/konis/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllRoles(r *gin.Context) {
	results := models.FindAllRoles()
	println(results)
	r.JSON(http.StatusOK, lib.Respont{
		Success: true,
		Message: "List All Roles",
		Result: results,
	})
}
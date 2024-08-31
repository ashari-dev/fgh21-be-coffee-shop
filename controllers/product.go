package controllers

import (
	"RGT/konis/lib"
	"RGT/konis/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)
func GetAllProduct(r *gin.Context) {
	results := models.FindAllProduct()
	println(results)
	r.JSON(http.StatusOK, lib.Respont{
		Success: true,
		Message: "List All Product",
		Result: results,
	})
}
func GEtOneProduct(r *gin.Context) {
	id,_ := strconv.Atoi(r.Param("id"))
	results := models.FindOneProduct(id)
	r.JSON(http.StatusOK, lib.Respont{
		Success: true,
		Message: "Id Product",
		Result: results,
	})
}
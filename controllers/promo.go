package controllers

import (
	"RGT/konis/lib"
	"RGT/konis/models"
	"RGT/konis/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)
func ListAllPromo(c *gin.Context) {
	promo, err := repository.GetAllPromo(models.Promo{})
	if err != nil {
		lib.HandlerNotfound(c, "Promo not found")
		return
	}

	lib.HandlerOK(c, "List All Promo", promo, nil)
}
func ListPromoById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	selectedPromo, err := repository.GetPromoById(id)

	if err != nil {
		lib.HandlerNotfound(c, "Data not found")
	}

	if selectedPromo.Id == 0 {
		lib.HandlerBadReq(c, "Failed to request the Promo")
		return
	}

	lib.HandlerOK(c, "Detail Promo", selectedPromo, nil)
}
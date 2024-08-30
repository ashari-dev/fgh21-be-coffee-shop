package controllers

import (
	"RGT/konis/lib"
	"RGT/konis/models"
	"RGT/konis/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)
func ListAllProductVariant(c *gin.Context) {
	product_variant, err := repository.GetAllProductVariant(models.ProductVariant{})
	if err != nil {
		lib.HandlerNotfound(c, "Product Variant not found")
		return
	}

	lib.HandlerOK(c, "List All Product Variant", product_variant, nil)
}
func ListProductVariantById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		lib.HandlerBadReq(c, "Invalid")
		return
	}

	dataVariant, err := repository.GetProductVariantById(id)

	if err != nil {
		lib.HandlerBadReq(c, "Product Variant not found")
		return
	}
	lib.HandlerOK(c, "Get Product Variant by Product Id", dataVariant, nil)
}
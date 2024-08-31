package controllers

import (
	"RGT/konis/dtos"
	"RGT/konis/lib"
	"RGT/konis/models"
	"RGT/konis/repository"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListAllProductOrderType (c *gin.Context) {
	productOrder, err := repository.FindAllProductOrderType(models.ProductOrderType{})

	if err != nil {
		lib.HandlerNotfound(c, "List Product Not Found")
		return
	}

	lib.HandlerOK(c, "List all product by order type", productOrder, nil)
}
func CreateNewProductOrderType (c *gin.Context) {

	var form dtos.ProductOrderType

	err := c.Bind(&form)
	fmt.Println(form)

	if err != nil {
		lib.HandlerBadReq(c, "Required to input data")
		return
	}

	productOrder, err := repository.AddNewProductOrderType(models.ProductOrderType{
		ProductId: form.ProductId,
		OrderTypeId: form.OrderTypeId,
	})

	if err != nil {
		fmt.Println(err)
		lib.HandlerBadReq(c, "Data not proper")
		return
	} 

	lib.HandlerOK(c, "Create product with order type success", productOrder, nil)
}
func DetailProductOrderTypeById (c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	selectedProduct, err := repository.FindProductOrderTypeById(id)

	if err != nil {
		lib.HandlerNotfound(c, "Data not found")
		return
	}

	if selectedProduct.Id == 0 {
		lib.HandlerBadReq(c, "Failed to request the product")
		return
	}

	lib.HandlerOK(c, "Detail product with order type", selectedProduct, nil)
}
func UpdateProductOrderType (c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var form dtos.ProductOrderType

	err := c.Bind(&form)

	if err != nil {
		lib.HandlerBadReq(c, "Required to input data")
		return
	}

	update, err := repository.EditProductOrderType(models.ProductOrderType{
		ProductId:   form.ProductId,
		OrderTypeId: form.OrderTypeId,
	}, id)

	if err != nil {
		lib.HandlerBadReq(c, "Required to input data")
		return
	}

	lib.HandlerOK(c, "Success edit product by order type", update, nil)
}
func DeleteProductOrderType (c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	selectedProduct, err := repository.FindProductOrderTypeById(id)

	if err != nil {
		lib.HandlerNotfound(c, "Data not found")
		return
	}

	err = repository.RemoveProductOrderType(models.ProductOrderType{}, id)

	if err != nil {
		lib.HandlerNotfound(c, "Data not found")
		return
	}

	lib.HandlerOK(c, "Delete the product", selectedProduct, nil)
}
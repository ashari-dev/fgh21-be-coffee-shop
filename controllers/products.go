package controllers

import (
	"RGT/konis/dtos"
	"RGT/konis/lib"
	"RGT/konis/models"
	"RGT/konis/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListAllProducts(c *gin.Context) {
	products, err := repository.GetAllProducts(models.Products{})
	if err != nil {
		lib.HandlerNotfound(c, "Products not found")
		return
	}

	lib.HandlerOK(c, "List All Products", products, nil)
}
func CreateProduct(c *gin.Context) {
	userId := 1
	var form dtos.Products

	err := c.Bind(&form)

	if err != nil {
		lib.HandlerBadReq(c, "Required to input data")
		return
	}

	newProduct, err := repository.AddNewProduct(models.Products{
		Title:       form.Title,
		Description: form.Description,
		Price:       form.Price,
		UserId:      &userId,
	})

	if err != nil {
		lib.HandlerBadReq(c, "Data not found")
		return
	}

	lib.HandlerOK(c, "Success to create new product", newProduct, nil)
}
func ListProductById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	selectedProduct, err := repository.GetProductById(id)

	if err != nil {
		lib.HandlerNotfound(c, "Data not found")
	}

	if selectedProduct.Id == 0 {
		lib.HandlerBadReq(c, "Failed to request the product")
		return
	}

	lib.HandlerOK(c, "Detail Product", selectedProduct, nil)
}
func UpdateProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var form dtos.Products

	err := c.Bind(&form)

	if err != nil {
		lib.HandlerBadReq(c, "Required to input data")
		return
	}

	update, err := repository.ChangeDataProduct(models.Products{
		Title:       form.Title,
		Description: form.Description,
		Price:       form.Price,
	}, id)

	if err != nil {
		lib.HandlerBadReq(c, "Required to input data")
		return
	}

	lib.HandlerOK(c, "Success Edit Product", update, nil)
}
func DeleteProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	selectUser, err := repository.GetProductById(id)

	if err != nil {
		lib.HandlerNotfound(c, "Data not found")
		return
	}

	err = repository.RemoveTheProduct(models.Products{}, id)

	if err != nil {
		lib.HandlerNotfound(c, "Data not found")
		return
	}

	lib.HandlerOK(c, "Delete the product", selectUser, nil)
}

func ListAllProductsSize(c *gin.Context) {
	products, err := repository.GetAllProductsSize(models.ProductsSizes{})
	if err != nil {
		lib.HandlerNotfound(c, "Products Sizes not found")
		return
	}

	lib.HandlerOK(c, "List All Products Sizes", products, nil)
}

func GetProductsById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		lib.HandlerBadReq(c, "Invalid")
		return
	}

	data, err := repository.FindProductSizeByProductId(id)

	if err != nil {
		lib.HandlerBadReq(c, "Product not found")
		return
	}
	lib.HandlerOK(c, "Get Product Sizes by Product Id", data, nil)
}

package controllers

import (
	"RGT/konis/dtos"
	"RGT/konis/lib"
	"RGT/konis/models"
	"RGT/konis/repository"
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	userId := c.GetInt("UserId")
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
		Stock:       form.Stock,
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
		return
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
		Stock:       form.Stock,
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

func ListAllCategoryProduct(c *gin.Context) {
	categoryproducts, err := repository.GetAllcategoryproduct(models.CategoryProduct{})
	if err != nil {
		lib.HandlerNotfound(c, "Category product not found")
		return
	}

	lib.HandlerOK(c, "List All Category Product", categoryproducts, nil)
}

func GetCategoryProductByCategoryId(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	category, err := repository.FindCategoryProductByCategoryId(id)

	if err != nil {
		lib.HandlerNotfound(c, "Data not found")
	}

	lib.HandlerOK(c, "Detail Category Product", category, nil)
}
func ListProductsWithPagination(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 100
	}
	products, err := repository.GetAllProductsWithPagination(page, limit)
	fmt.Println(err)
	if err != nil {
		lib.HandlerNotfound(c, "Products not found")
		return
	}
	log.Println(products)

	lib.HandlerOK(c, "List All Products", products, nil)
}

func ListAllOurProductsWithPagination(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 100
	}
	products, err := repository.GetAllOurProductsWithPagination(page, limit)
	fmt.Println(err)
	if err != nil {
		lib.HandlerNotfound(c, "Products not found")
		return
	}

	lib.HandlerOK(c, "List All Products", products, nil)
}

// func ListAllFilterProductsWithPagination(c *gin.Context) {
// 	data := c.Query("data")
// 	page, _ := strconv.Atoi(c.Query("page"))
// 	limit, _ := strconv.Atoi(c.Query("limit"))
// 	if page < 1 {
// 		page = 1
// 	}
// 	if limit < 1 {
// 		limit = 100
// 	}
// 	products, err := repository.GetAllProductsWithFilterPagination(data, page, limit)
// 	fmt.Println(err)
// 	if err != nil {
// 		lib.HandlerNotfound(c, "Products not found")
// 		return
// 	}
// 	log.Println(products)

// 	lib.HandlerOK(c, "List Filter Products", products, nil)
// }

func ListAllFilterProductsWithPagination(c *gin.Context) {
	title := c.Query("title")
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 100
	}
	products, err := repository.GetAllProductsWithFilterPagination(title, page, limit)
	fmt.Println(err)
	if err != nil {
		lib.HandlerNotfound(c, "Products not found")
		return
	}
	log.Println(products)

	lib.HandlerOK(c, "List Filter Products", products, nil)
}

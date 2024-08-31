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

func ListAllCarts(ctx *gin.Context) {
	id := ctx.GetInt("UserId")
	carts, err := repository.FindAllCarts(id)

	if err != nil {
		fmt.Println(err)
		lib.HandlerNotfound(ctx, "Carts Not Found")
		return
	}

	lib.HandlerOK(ctx, "List All Carts", carts, nil)
}
func CreateOneCarts(ctx *gin.Context) {
	id := ctx.GetInt("UserId")

	// id := 1

	form := dtos.FormCarts{}
	err := ctx.Bind(&form)

	if err != nil {
		fmt.Println(err)
		lib.HandlerBadReq(ctx, "Quantity is null")
		return
	}

	fmt.Println(id)

	carts, err := repository.CreateCarts(models.Carts{
		Quantity:     form.Quantity,
		VariantId:    form.VariantProduct,
		SizesProduct: form.SizesProduct,
		ProductId:    form.ProductId,
		UserId:       id,
	})

	if err != nil {
		lib.HandlerBadReq(ctx, "Failed add to carts")
		return
	}

	lib.HandlerOK(ctx, "Add to Carts success", carts, nil)
}

func DeleteOneCarts(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	selectUser, err := repository.GetCartsById(id)

	if err != nil {
		lib.HandlerNotfound(ctx, "Data not found")
		return
	}

	err = repository.DeleteCarts(models.Carts{}, id)

	if err != nil {
		lib.HandlerNotfound(ctx, "Data not found")
		return
	}

	lib.HandlerOK(ctx, "Delete the product", selectUser, nil)
}

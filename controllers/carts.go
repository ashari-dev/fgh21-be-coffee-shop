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
	userId := ctx.GetInt("UserId")
	productId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		lib.HandlerBadReq(ctx, "Invalid Product Id")
		return
	}

	form := dtos.FormCarts{}
	err = ctx.Bind(&form)

	if err != nil {
		fmt.Println(err)
		lib.HandlerBadReq(ctx, "Invalid Input Data")
		return
	}

	carts, err := repository.CreateCarts(models.Carts{
		TransactionDetail: form.TransactionDetail,
		Quantity:          form.Quantity,
		VariantId:         form.Variant,
		ProductSizeId:     form.ProductSize,
		ProductId:         productId,
		UserId:            userId,
	})

	if err != nil {
		lib.HandlerBadReq(ctx, "Failed add to carts")
		return
	}

	lib.HandlerOK(ctx, "Add to Carts success", carts, nil)
}

func DeleteCartsByUserId(ctx *gin.Context) {
	id := ctx.GetInt("UserId")
	// fmt.Println(id)
	// selectUser, err := repository.GetCartsByUserId(id)

	// if err != nil {
	// 	lib.HandlerNotfound(ctx, "Carts not found")
	// 	return
	// }

	err := repository.DeleteCarts(models.Carts{}, id)
	fmt.Println(err)
	if err != nil {
		lib.HandlerNotfound(ctx, "Delete Failed")
		return
	}

	lib.HandlerOK(ctx, "Delete the product", id, nil)
}

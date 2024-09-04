package controllers

import (
	"RGT/konis/dtos"
	"RGT/konis/lib"
	"RGT/konis/models"
	"RGT/konis/repository"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetTransactionDetailById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		lib.HandlerBadReq(c, "Invalid")
		return
	}

	data, err := repository.FindTransactionDetailById(id)

	fmt.Println(err)

	if err != nil {
		lib.HandlerBadReq(c, "Data not found")
		return
	}
	lib.HandlerOK(c, "Get Transaction Detail by Id", data, nil)
}

func GetTransactionProductById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		lib.HandlerBadReq(c, "Invalid")
		return
	}

	data, err := repository.FindTransactionProductById(id)

	fmt.Println(err)

	if err != nil {
		lib.HandlerBadReq(c, "Data not found")
		return
	}
	lib.HandlerOK(c, "Get Transaction Detail by Id", data, nil)
}

func GetAllTransactionByUserId(c *gin.Context) {
	id := c.GetInt("UserId")
	fmt.Println(id)
	data, err := repository.FindTransactionByUserId(id)

	fmt.Println(err)

	if err != nil {
		lib.HandlerBadReq(c, "Transaction not found")
		return
	}
	lib.HandlerOK(c, "Get Transaction by User Id", data, nil)
}

func CreateTransactionDetails(c *gin.Context) {
	var formTransaction dtos.TransactionDetail
	err := c.Bind(&formTransaction)
	if err != nil {
		lib.HandlerBadReq(c, "Invalid")
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		lib.HandlerBadReq(c, "Invalid Product Id")
		return
	}
	data, err := repository.CreateTransactionDetail(models.TransactionDetail{
		Quantity:      formTransaction.Quantity,
		ProductId:     id,
		VariantId:     formTransaction.Variant,
		ProductSizeId: formTransaction.ProductSize,
	})

	fmt.Println(err)

	if err != nil {
		lib.HandlerBadReq(c, "Failed To Created Transaction")
		return
	}
	lib.HandlerOK(c, "Create Transaction Success", data, nil)
}

func CreateTransaction(c *gin.Context) {
	var formTransaction dtos.FormTransaction
	err := c.Bind(&formTransaction)
	userId := c.GetInt("UserId")
	fmt.Println(err)
	if err != nil {
		lib.HandlerBadReq(c, "Invalid")
		return
	}
	fmt.Println(userId)
	noOrder := rand.Intn(90000) + 10000
	for i := range formTransaction.TransactionDetail {
		repository.CreateTransaction(models.Transaction{
			NoOrder:             noOrder,
			AddFullName:         formTransaction.FullName,
			AddEmail:            formTransaction.Email,
			AddAddress:          formTransaction.Address,
			Payment:             formTransaction.Payment,
			UserId:              userId,
			TransactionDetail:   formTransaction.TransactionDetail[i],
			OrderTypeId:         formTransaction.OrderType,
			TransactionStatusId: formTransaction.TransactionStatus,
		})
	}
	// data, err := repository.CreateTransaction(models.Transaction{
	// 	NoOrder:             noOrder,
	// 	AddFullName:         formTransaction.FullName,
	// 	AddEmail:            formTransaction.Email,
	// 	AddAddress:          formTransaction.Address,
	// 	Payment:             formTransaction.Payment,
	// 	UserId:              userId,
	// 	TransactionDetail:   formTransaction.TransactionDetail,
	// 	OrderTypeId:         formTransaction.OrderType,
	// 	TransactionStatusId: formTransaction.TransactionStatus,
	// })
	fmt.Println(err)
	if err != nil {
		lib.HandlerBadReq(c, "Invalid Data")
		return
	}

	lib.HandlerOK(c, "transaction success", nil, nil)
}

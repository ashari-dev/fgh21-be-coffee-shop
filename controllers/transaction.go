package controllers

import (
	"RGT/konis/dtos"
	"RGT/konis/lib"
	"RGT/konis/models"
	"RGT/konis/repository"
	"fmt"
	"math"
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
	fmt.Println(err)
	if err != nil {
		lib.HandlerBadReq(c, "Invalid Data")
		return
	}

	lib.HandlerOK(c, "transaction success", nil, nil)
}


func GetALLTransactions(c *gin.Context) {
	search := c.Query("search")
	limitParam := c.Query("limit")
	limit, _ := strconv.Atoi(limitParam)
	pageParam := c.Query("page")
	page, _ := strconv.Atoi(pageParam)

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	transaction, count := repository.FindAllTransactions(search, page, limit)
	totalPage := math.Ceil(float64(count) / float64(limit))

	next := 0
	if int(totalPage) >= 1 {
		next = int(totalPage) - page
	}
	prev := page
	if page >= 1 {
		prev = page - 1
	}

	lib.HandlerOK(c, "List All transactions", transaction, lib.PageInfo{
		TotalData: count,
		TotalPage: int(totalPage),
		Page:      page,
		Limit:     limit,
		Next:      &next,
		Prev:      &prev,
	})
}

func GetALLTransactionsByStatusId(c *gin.Context) {
	search := c.Query("search")
	searchId, _ := strconv.Atoi(search)

	limitParam := c.Query("limit")
	limit, _ := strconv.Atoi(limitParam)
	pageParam := c.Query("page")
	page, _ := strconv.Atoi(pageParam)

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 3
	}

	transaction, count := repository.FindTransactionsByStatusId(searchId, page, limit)
	totalPage := math.Ceil(float64(count) / float64(limit))

	next := 0
	if int(totalPage) >= 1 {
		next = int(totalPage) - page
	}
	prev := page
	if page >= 1 {
		prev = page - 1
	}

	lib.HandlerOK(c, "List All transactions", transaction, lib.PageInfo{
		TotalData: count,
		TotalPage: int(totalPage),
		Page:      page,
		Limit:     limit,
		Next:      &next,
		Prev:      &prev,
	})
}

func UpdateTransactionStatus(c *gin.Context) {
	noOrder, _ := strconv.Atoi(c.Param("id"))
	var form dtos.FormTransaction

	err := c.Bind(&form)
	if err != nil {
		lib.HandlerBadReq(c, "Required to input data")
		return
	}
	update, err := repository.EditTransactionStatus(models.Transaction{
		TransactionStatusId: form.TransactionStatus,
		}, noOrder)
		
		if err != nil {
		fmt.Println(err)
		lib.HandlerBadReq(c, "Required to input data")
		return
	}

	lib.HandlerOK(c, "Success edit product by order type", update, nil)
}

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

// SELECT no_order, products.title, transaction_details.quantity, product_variants.name, product_sizes.name, order_types.name, products.price
// FROM transactions
// INNER JOIN transaction_details on transactions.transaction_detail_id = transaction_details.id
// INNER JOIN products on transaction_details.product_id = products.id
// INNER JOIN product_sizes on transaction_details.product_size_id = product_sizes.id
// INNER JOIN product_variants on transaction_details.variant_id = product_variants.id
// INNER JOIN order_types on transactions.order_type_id = order_types.id

func GetAllTransactionByUserId(c *gin.Context) {
	id := c.GetInt("UserId")
	fmt.Println(id)
	data, err := repository.FindTransactionByUserId(id)

	fmt.Println(err)

	if err != nil {
		lib.HandlerBadReq(c, "Transaction not found")
		return
	}
	lib.HandlerOK(c, "Get Transaction Detail by User Id", data, nil)
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
	data, err := repository.CreateTransaction(models.Transaction{
		NoOrder:             noOrder,
		AddFullName:         formTransaction.FullName,
		AddEmail:            formTransaction.Email,
		AddAddress:          formTransaction.Address,
		Payment:             formTransaction.Payment,
		UserId:              userId,
		TransactionDetail:   formTransaction.TransactionDetail,
		OrderTypeId:         formTransaction.OrderType,
		TransactionStatusId: formTransaction.TransactionStatus,
	})
	fmt.Println(err)
	if err != nil {
		lib.HandlerBadReq(c, "Invalid Data")
		return
	}

	lib.HandlerOK(c, "transaction success", data, nil)
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
		limit = 3
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

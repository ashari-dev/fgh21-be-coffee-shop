package dtos

type ProductOrderType struct {
	ProductId 	int `form:"productId"`
	OrderTypeId int `form:"orderTypeId"`
}
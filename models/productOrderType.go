package models

type ProductOrderType struct {
	Id 			int `json:"id"`
	ProductId 	int `json:"productId" db:"product_id"`
	OrderTypeId int `json:"orderTypeId" db:"order_type_id"`
}
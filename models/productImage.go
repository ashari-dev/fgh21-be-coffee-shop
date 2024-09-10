package models

type ProductImage struct {
	Id int `json:"id"`
	Image string `json:"image"`
	ProductId int `json:"productId" db:"product_id"`
}
package models

type ProductVariant struct{
	Id int `json:"id"`
	Name string `json:"product_variant" db:"name"`
	Add_price int `json:"add_price" db:"add_price"`
	Stock int `json:"stock" db:"stock"`
	Product_id int `json:"product_id" db:"product_id"`
}
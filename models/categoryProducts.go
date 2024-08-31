package models

type CategoryProduct struct {
	Id          int `json:"id"`
	Category_id int `json:"category_id" db:"category_id"`
	Product_id  int `json:"product_id" db:"product_id"`
}

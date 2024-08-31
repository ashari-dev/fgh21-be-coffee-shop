package models

type ProductsSizes struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Add_price  string `json:"add_price" db:"add_price"`
	Product_id *int   `json:"product_id" db:"product_id"`
}

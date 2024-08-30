package models

type Products struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	UserId      *int   `json:"userId" db:"user_id"`
}
type ProductsSizes struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Add_price  string `json:"add_price" db:"add_price"`
	Product_id *int   `json:"product_id" db:"product_id"`
}

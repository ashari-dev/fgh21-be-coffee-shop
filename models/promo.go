package models

type Promo struct {
	Id int `json:"id"`
	Name string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Discount int `json:"discount" db:"discount"`
}
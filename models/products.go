package models

type Products struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Price int `json:"price"`
	UserId *int `json:"userId" db:"user_id"`
}
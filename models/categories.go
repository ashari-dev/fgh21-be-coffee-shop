package models

type Categories struct {
	Id   int    `json:"id"`
	Name string `json:"name" db:"name"`
}

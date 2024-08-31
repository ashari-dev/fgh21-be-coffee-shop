package models

type Roles struct {
	Id   int    `json:"id"`
	Name string `json:"name" db:"name"`
}

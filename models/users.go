package models

type Users struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"-"`
	UserId   int    `json:"userIs" db:"user_id"`
}

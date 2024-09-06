package models

type Users struct {
	Id       int    `json:"id"`
	Email    string `json:"email" form:"email"`
	Password string `json:"-" form:"password"`
	RoleId   int    `json:"roleId" db:"role_id"`
}

type InsertUsers struct {
	Id       int    `json:"id"`
	Email    string `json:"email" db:"email"`
	Password string `json:"-" form:"password"`
	RoleId   int    `json:"role_id" db:"role_id"`
}

type InsertProfile struct {
	Id          int     `json:"id"`
	FullName    string  `json:"full_name" db:"full_name"`
	PhoneNumber *string `json:"phone_number" db:"phone_number"`
	Address     *string `json:"address" db:"address"`
	Image       *string `json:"image" db:"image"`
	UserId      int     `json:"user_id" db:"user_id"`
}

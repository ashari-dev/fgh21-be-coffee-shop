package dtos

type Profile struct {
	Id          int     `json:"id" db:"id"`
	FullName    string  `json:"fullName" form:"fullName" db:"full_name"`
	PhoneNumber *string `json:"phoneNumber" form:"phoneNumber" db:"phone_number"`
	Address     string  `json:"address" form:"address"`
	Image       *string `json:"image" db:"image"`
	UserId      int     `json:"userId" db:"user_id"`
}

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required,min=8"`
	RoleId   int    `json:"roleId"`
}

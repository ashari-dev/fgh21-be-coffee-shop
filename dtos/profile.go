package dtos

type FormProfile struct {
	Id          int     `json:"id" db:"id"`
	FullName    string  `json:"fullName" form:"fullName" db:"full_name"`
	PhoneNumber *string `json:"phoneNumber" form:"phoneNumber" db:"phone_number"`
	Address     string  `json:"address" form:"address"`
	Image       *string `json:"image" db:"image"`
	UserId      int     `json:"userId" db:"user_id"`
}


package models

type Profile struct {
	Id          int     `json:"id"`
	FullName    string  `json:"fullName" db:"full_name"`
	PhoneNumber *string `json:"phoneNumber" db:"phone_number"`
	Address     *string `json:"address"`
	Image       *string `json:"image"`
	UserId      int     `json:"userId" db:"user_id"`
}

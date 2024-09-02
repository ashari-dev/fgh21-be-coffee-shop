package models

type Profile struct {
	Id          int     `json:"id"`
	FullName    string  `json:"fullName" db:"full_name"`
	PhoneNumber *string `json:"phoneNumber" db:"phone_number"`
	Address     *string `json:"address"`
	Image       *string `json:"image"`
	UserId      int     `json:"userId" db:"user_id"`
}

type ProfileJoinUser struct {
	Id          int     `json:"id"`
	FullName    string  `json:"fullName" db:"full_name"`
	PhoneNumber *string `json:"phoneNumber" db:"phone_number"`
	Email       string  `json:"email"`
	// Password    *string `json:"-" form:"password"`
	Address *string `json:"address"`
	Image   *string `json:"image"`
	// RoleId      *int    `json:"roleId" db:"role_id"`
}

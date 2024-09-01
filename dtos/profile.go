package dtos

type FormProfile struct {
	Id          int     `json:"id" db:"id"`
	FullName    string  `json:"fullName" form:"fullName" db:"full_name"`
	PhoneNumber *string `json:"phoneNumber" form:"phoneNumber" db:"phone_number"`
	Address     string  `json:"address" form:"address"`
	Image       *string `json:"image" db:"image"`
	UserId      int     `json:"userId" db:"user_id"`
}

type ProfileJoinUser struct {
	Id          int     `json:"id"`
	FullName    string  `json:"fullName" form:"fullName" db:"full_name"`
	Email       string  `json:"email" form:"email"`
	Password    *string  `json:"-" form:"password"`
	PhoneNumber string  `json:"phoneNumber" form:"phoneNumber" db:"phone_number"`
	Address     string  `json:"address" form:"address"`
	Image       *string `json:"image"`
	RoleId      *int    `json:"roleId" db:"role_id"`
}

type ProfileUser struct {
	Id          int     `json:"id"`
	FullName    string  `json:"fullName" form:"fullName" db:"full_name"`
	Email       string  `json:"email" form:"email"`
	PhoneNumber *string  `json:"phoneNumber" form:"phoneNumber" db:"phone_number"`
	Address     *string  `json:"address" form:"address"`
	Image       *string `json:"image"`
	RoleId      int    `json:"roleId" db:"role_id"`
}


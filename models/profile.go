package models

import "RGT/konis/dtos"

type JoinProfile struct {
	Id       int    `json:"id"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"-" form:"password" binding:"required"`
	RoleId   int    `json:"roleId" db:"role_id"`
	Profile  dtos.Profile
}

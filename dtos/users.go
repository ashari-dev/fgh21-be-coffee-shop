package dtos

import "mime/multipart"

type FormUser struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required,min=6"`
	RoleId   int    `form:"roleId"`
}

type CreateUserProfileInput struct {
	Email       string                `form:"email" binding:"required"`
	Password    string                `form:"password" binding:"required,min=6"`
	RoleId      int                   `form:"roleId"`
	FullName    string                `form:"fullName"`
	PhoneNumber string                `form:"phoneNumber"`
	Address     string                `form:"address"`
	Image       *multipart.FileHeader `form:"image"`
}

package dtos

type FormUser struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required,min=6"`
	RoleId   int    `form:"roleId"`
}

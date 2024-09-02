package dtos

type LoginForm struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

type RegisterForm struct {
	FullName        string `form:"fullName" binding:"required"`
	Email           string `form:"email" binding:"email"`
	Password        string `form:"password" binding:"min=8"`
	ConfirmPassword string `form:"cPassword" binding:"eqfield=Password"`
}

type Token struct {
	Token string `json:"token"`
}

package controllers

import (
	"RGT/konis/dtos"
	"RGT/konis/lib"
	"RGT/konis/models"
	"RGT/konis/repository"

	"github.com/gin-gonic/gin"
)


func AuthLogin(c *gin.Context) {
	formLogin := dtos.LoginForm{}
	err := c.Bind(&formLogin)

	if err != nil {
		lib.HandlerBadReq(c, "email and password is null")
		return
	}

	found, err := repository.FindUserByEmail(formLogin.Email)
	if err != nil {
		lib.HandlerBadReq(c, "Wrong email and password")
		return
	}
	if found == (models.Users{}) {
		lib.HandlerUnauthorized(c, "Wrong email")
		return
	}

	isVerified := lib.Verify(formLogin.Password, found.Password)

	if !isVerified {
		lib.HandlerUnauthorized(c, "Wrong password")
		return
	} else {
		token := lib.GenerateUserTokenById(found.Id)
		lib.HandlerOK(c, "Login success", dtos.Token{Token: token}, nil)
	}
}

func AuthRegister(c *gin.Context) {
form := models.JoinProfile{}
	user := dtos.User{}

	err := ctx.Bind(&form)
	if err != nil {
		lib.HandlerBadReq(ctx, "Register Failed")
		return
	}

	roleId := 1

	repository.CreateProfile(form, roleId)

	user.Email = form.Email
	user.Password = form.Password
	createUser := repository.CreateUser(user, roleId)

	lib.HandlerOK(ctx, "Register success", createUser, lib.PageInfo{})

}

package controllers

import (
	"RGT/konis/dtos"
	"RGT/konis/lib"
	"RGT/konis/models"
	"RGT/konis/repository"
	"fmt"

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
	formRegister := dtos.RegisterForm{}
	err := c.Bind(&formRegister)
	fmt.Println(formRegister)
	if err != nil {
		lib.HandlerBadReq(c, "format invalid")
		return
	}

	user, err := repository.CreateUser(models.Users{
		Email:    formRegister.Email,
		Password: formRegister.Password,
		RoleId:   1,
	})
	if err != nil {
		lib.HandlerBadReq(c, "data not verified")
		return
	}

	profile, err := repository.CreateProfile(models.Profile{
		FullName: formRegister.FullName,
		UserId:   user.Id,
	})
	if err != nil {
		lib.HandlerBadReq(c, "data not verified")
		return
	}

	lib.HandlerOK(c, "Register success", profile, nil)
}

// pull

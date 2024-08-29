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

func RegisterLogin(c *gin.Context) {

}

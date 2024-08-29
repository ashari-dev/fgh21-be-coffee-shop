package controllers

import (
	"RGT/konis/dtos"
	"RGT/konis/lib"
	"RGT/konis/models"
	"RGT/konis/repository"
	"fmt"

	"github.com/gin-gonic/gin"
)

func AuthRegister(ctx *gin.Context) {
	form := models.JoinProfile{}
	user := dtos.User{}

	err := ctx.Bind(&form)
	fmt.Println(form)
	if err != nil {
		lib.HandlerBadReq(ctx, "Register Failed")
		return
	}

	// dataUser := models.FindOneUserByEmail(form.Email)

	// if dataUser.Email == form.Email {
	// 	ctx.JSON(http.StatusBadRequest, lib.Response{
	// 		Success: false,
	// 		Message: "Email already exist",
	// 	})
	// 	return
	// }

	roleId := 1

	repository.CreateProfile(form, roleId)

	user.Email = form.Email
	user.Password = form.Password
	createUser := repository.CreateUser(user, roleId)

	lib.HandlerOK(ctx, "Register success", createUser, lib.PageInfo{})
}

package controllers

import (
	"RGT/konis/dtos"
	"RGT/konis/lib"
	"RGT/konis/models"
	"RGT/konis/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	data, err := repository.FindAllUsers()

	if err != nil {
		lib.HandlerBadReq(c, "Invalid")
		return
	}
	lib.HandlerOK(c, "List all users", data, nil)
}

func GetUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		lib.HandlerBadReq(c, "Invalid")
		return
	}

	data, err := repository.FindUserById(id)

	if err != nil {
		lib.HandlerBadReq(c, "Data not found")
		return
	}
	lib.HandlerOK(c, "Get user by id", data, nil)
}

func CreateUser(c *gin.Context) {
	formUser := dtos.FormUser{}
	err := c.Bind(&formUser)

	if err != nil {
		lib.HandlerBadReq(c, "data not verified")
		return
	}

	roleId := 1

	data, err := repository.CreateUser(models.Users{
		Email:    formUser.Email,
		Password: formUser.Password,
		RoleId:   roleId,
	})

	if err != nil {
		lib.HandlerBadReq(c, "data not verified")
		return
	}

	lib.HandlerOK(c, "Create user success", data, nil)
}

func UpdateUserById(c *gin.Context) {
	formUser := dtos.FormUser{}
	err := c.Bind(&formUser)
	if err != nil {
		lib.HandlerBadReq(c, "data not verified")
		return
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		lib.HandlerBadReq(c, "data not verified")
		return
	}

	roleId := 1

	data, err := repository.UpdateUserById(models.Users{
		Email:    formUser.Email,
		Password: formUser.Password,
		RoleId:   roleId,
	}, id)

	if data.Id == 0 {
		lib.HandlerNotfound(c, "Data not found")
		return
	}

	if err != nil {
		lib.HandlerBadReq(c, "data not verified")
		return
	}

	lib.HandlerOK(c, "Update user success", data, nil)
}

func DeleteUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		lib.HandlerBadReq(c, "Invalid")
		return
	}

	data, err := repository.DeleteUserById(id)

	if data.Id == 0 {
		lib.HandlerNotfound(c, "Data not found")
		return
	}

	if err != nil {
		lib.HandlerBadReq(c, "Data not found")
		return
	}
	lib.HandlerOK(c, "Delete user success", data, nil)
}

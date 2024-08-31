package controllers

import (
	"RGT/konis/dtos"
	"RGT/konis/lib"
	"RGT/konis/models"
	"RGT/konis/repository"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FindProfileById(c *gin.Context) {
	id := c.GetInt("UserId")
	profile, err := repository.FindProfileById(id)
	fmt.Println(id)

	if err != nil {
		lib.HandlerBadReq(c, "Profile not found ")
		return
	}

	lib.HandlerOK(c, "Success Find Profile By UserId", profile, nil)
}

func UpdateProfile(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	// id := c.GetInt("userId")
	form := dtos.ProfileJoinUser{}
	fmt.Println(form)

	err := c.Bind(&form)

	if err != nil {
		lib.HandlerBadReq(c, "Required to input data")
		return
	}

	repository.UpdateUserById(models.Users{
		Email:    form.Email,
		Password: *form.Password,
	}, id)

	updateProfile, err := repository.UpdateProfile(models.Profile{
		FullName:    form.FullName,
		PhoneNumber: &form.PhoneNumber,
		Address:     &form.Address,
	}, id)

	if err != nil {
		lib.HandlerBadReq(c, "Required to input data")
		return
	}

	lib.HandlerOK(c, "Success Edit Product", updateProfile, nil)
}

func DeleteProfile(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	repository.RemoveProfile(id)
	selectUser, err := repository.DeleteUserById(id)

	if err != nil {
		lib.HandlerNotfound(c, "Data not found")
		return
	}

	lib.HandlerOK(c, "Delete the product", selectUser, nil)
}

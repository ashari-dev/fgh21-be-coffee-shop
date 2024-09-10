package controllers

import (
	"RGT/konis/dtos"
	"RGT/konis/lib"
	"RGT/konis/models"
	"RGT/konis/repository"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

func CreateUserWithProfile(c *gin.Context) {
	var input dtos.CreateUserProfileInput

	if err := c.Bind(&input); err != nil {
		lib.HandlerBadReq(c, "Invalid input data")
		return
	}

	file, err := c.FormFile("profileImage")
	if err != nil {
		lib.HandlerBadReq(c, "Image upload failed")
		return
	}

	allowedExtensions := map[string]bool{".jpg": true, ".jpeg": true, ".png": true}
	fileExt := strings.ToLower(filepath.Ext(file.Filename))
	if !allowedExtensions[fileExt] {
		lib.HandlerBadReq(c, "Invalid file extension. Allowed: .jpg, .jpeg, .png")
		return
	}

	newFileName := uuid.New().String() + fileExt
	uploadDir := "./img/profile/"
	fullFilePath := uploadDir + newFileName

	if err := c.SaveUploadedFile(file, fullFilePath); err != nil {
		lib.HandlerBadReq(c, "Failed to upload image")
		return
	}

	encryptedPassword := lib.Encrypt(input.Password)

	user, err := repository.CreateinsertUser(models.InsertUsers{
		Email:    input.Email,
		Password: encryptedPassword,
		RoleId:   input.RoleId,
	})
	if err != nil {
		lib.HandlerBadReq(c, "Failed to create user: "+err.Error())
		return
	}

	imageURL := "http://localhost:8000/img/profile/" + newFileName

	profile, err := repository.CreateinsertProfile(models.InsertProfile{
		FullName:    input.FullName,
		PhoneNumber: &input.PhoneNumber,
		Address:     &input.Address,
		Image:       &imageURL,
		UserId:      user.Id,
	})
	if err != nil {
		lib.HandlerBadReq(c, "Failed to create profile: "+err.Error())
		return
	}

	lib.HandlerOK(c, "User and Profile created successfully", profile, nil)
}

package controllers

import (
	"RGT/konis/dtos"
	"RGT/konis/lib"
	"RGT/konis/models"
	"RGT/konis/repository"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetALLProfiles(c *gin.Context) {
	profile, err := repository.FindAllProfiles()

	if err != nil {
		lib.HandlerBadReq(c, "Failed to get all profile")
		return
	}

	lib.HandlerOK(c, "List All Category", profile, nil)
}

func CreateProfileJoinUser(c *gin.Context) {
	inputUser := dtos.FormProfileJoinUser{}
	err := c.Bind(&inputUser)
	fmt.Println(inputUser)
	if err != nil {
		lib.HandlerBadReq(c, "format invalid")
		return
	}

	user, err := repository.CreateUser(models.Users{
		Email:    inputUser.Email,
		Password: inputUser.Password,
		RoleId:   inputUser.RoleId,
	})
	if err != nil {
		lib.HandlerBadReq(c, "data not verified")
		return
	}

	profile, err := repository.CreateProfileJoinUser(models.Profile{
		FullName:    inputUser.FullName,
		PhoneNumber: &inputUser.PhoneNumber,
		Address:     &inputUser.Address,
		// Image:       &inputUser.Image,
		UserId: user.Id,
	})
	if err != nil {
		lib.HandlerBadReq(c, "data not verified")
		return
	}

	lib.HandlerOK(c, "Register success", profile, nil)
}

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
	// id, _ := strconv.Atoi(c.Param("id"))
	id := c.GetInt("UserId")
	form := dtos.ProfileJoinUser{}
	// fmt.Println(form)

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

func UploadProfileImage(c *gin.Context) {
	id := c.GetInt("UserId")
	maxFile := 500 * 1024
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, int64(maxFile))

	file, err := c.FormFile("profileImg")
	if err != nil {
		if err.Error() == "http: request body too large" {
			lib.HandlerMaxFile(c, "file size too large, max capacity 500 kb")
			return
		}
		lib.HandlerBadReq(c, "not file to upload")
		return
	}

	allowExt := map[string]bool{".jpg": true, ".jpeg": true, ".png": true}
	fileExt := strings.ToLower(filepath.Ext(file.Filename))
	if !allowExt[fileExt] {
		lib.HandlerBadReq(c, "extension file not validate")
		return
	}

	newFile := uuid.New().String() + fileExt

	uploadDir := "./img/profile/"
	if err := c.SaveUploadedFile(file, uploadDir+newFile); err != nil {
		lib.HandlerBadReq(c, "upload failed")
		return
	}

	tes := "http://localhost:8000/img/profile/" + newFile

	fmt.Println(tes)
	fmt.Println(id)
	profile, err := repository.UpdateProfileImage(models.Profile{Image: &tes}, id)
	if err != nil {
		lib.HandlerBadReq(c, "upload failed")
		return
	}

	lib.HandlerOK(c, "Upload success", profile, nil)
}

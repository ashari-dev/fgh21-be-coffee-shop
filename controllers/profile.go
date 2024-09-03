package controllers

import (
	"RGT/konis/dtos"
	"RGT/konis/lib"
	"RGT/konis/models"
	"RGT/konis/repository"
	"fmt"
	"math"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetALLProfiles(c *gin.Context) {
	search := c.Query("search")
	limitParam := c.Query("limit")
	limit, _ := strconv.Atoi(limitParam)
	pageParam := c.Query("page")
	page, _ := strconv.Atoi(pageParam)

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 3
	}

	profile, count := repository.FindAllProfiles(search, page, limit)
	totalPage := math.Ceil(float64(count) / float64(limit))

	next := 0
	if int(totalPage) >= 1 {
		next = int(totalPage) - page
	}
	prev := page
	if page >= 1 {
		prev = page - 1
	}

	lib.HandlerOK(c, "List All Profiles", profile, lib.PageInfo{
		TotalData: count,
		TotalPage: int(totalPage),
		Page:      page,
		Limit:     limit,
		Next:      &next,
		Prev:      &prev,
	})
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
		Image:       &inputUser.Image,
		UserId:      user.Id,
	})
	if err != nil {
		lib.HandlerBadReq(c, "data not verified")
		return
	}

	lib.HandlerOK(c, "Register success", profile, nil)
}

func FindProfileById(c *gin.Context) {
	id := c.GetInt("UserId")
	if id == 0 {
		id, _ = strconv.Atoi(c.Param("id"))
	}

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
	if id == 0 {
		id = c.GetInt("UserId")
	}
	form := dtos.ProfileJoinUser{}

	err := c.Bind(&form)

	if err != nil {
		lib.HandlerBadReq(c, "Required to input data")
		return
	}

	if form.Password == nil {
		emptyPassword := ""
		form.Password = &emptyPassword
	}

	user, err := repository.UpdateUserById(models.Users{
		Email:    form.Email,
		Password: *form.Password,
	}, id)

	if err != nil {
		lib.HandlerBadReq(c, "Cannot update user")
		return
	}

	updateProfile, err := repository.UpdateProfile(models.Profile{
		FullName:    form.FullName,
		PhoneNumber: &form.PhoneNumber,
		Address:     &form.Address,
	}, id)

	if err != nil {
		lib.HandlerBadReq(c, "Required to input data")
		return
	}

	updateProfile.Email = user.Email

	lib.HandlerOK(c, "Success update profile", updateProfile, nil)
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
	fmt.Println(id)

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
	if id == 0 {
		lib.HandlerBadReq(c, "User not found")
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

	profile, err := repository.UpdateProfileImage(models.Profile{Image: &tes}, id)
	if err != nil {
		lib.HandlerBadReq(c, "upload failed")
		return
	}

	lib.HandlerOK(c, "Upload success", profile, nil)
}

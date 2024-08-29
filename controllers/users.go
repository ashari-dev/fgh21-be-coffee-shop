package controllers

import (
	"RGT/konis/lib"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	lib.HandlerOK(c, "berhasil", nil, nil)
}

package controllers

import (
	"RGT/konis/lib"
	"RGT/konis/repository"

	"github.com/gin-gonic/gin"
)

func ListAllProfiles(ctx *gin.Context) {
	profiles := repository.FindAllProfiles()
	lib.HandlerOK(ctx, "List all profiles", profiles, lib.PageInfo{})
}

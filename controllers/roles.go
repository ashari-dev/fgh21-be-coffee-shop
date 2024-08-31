package controllers

import (
	"RGT/konis/lib"
	"RGT/konis/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllRoles(c *gin.Context) {
	results := repository.FindAllRoles()
	println(results)
	lib.HandlerOK(c, "List All Roles", results, nil)

}

func GetOneRoles(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		lib.HandlerBadReq(c, "Inavlid id")
		return
	}

	results := repository.FindOneRoles(id)
	if results.Id == 0 {
		lib.HandlerNotfound(c, "id is not found")
		return
	}

	lib.HandlerOK(c, "List All Roles", results, nil)

}

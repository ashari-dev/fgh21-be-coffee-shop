package controllers

import (
	"RGT/konis/lib"
	"RGT/konis/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SeeAllOrderTypes(c *gin.Context) {
	results, err := repository.FindAllOrderTypes()
	if err != nil {
		lib.HandlerNotfound(c, "Data Not Found")
	}
	lib.HandlerOK(c, "See All Order Types", results, nil)
}
func SeeOrderTypesById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	orderTypes, err := repository.FindOneOrderTypes(id)

	if err != nil {
		lib.HandlerNotfound(c, "Data not found")
	}

	if orderTypes.Id == 0 {
		lib.HandlerBadReq(c, "Failed to request")
		return
	}

	lib.HandlerOK(c, "Detail Order Types", orderTypes, nil)
}

package controllers

import (
	"RGT/konis/lib"
	"RGT/konis/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetTRansactionStatusById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	transactionStatus, err := repository.FindTransactionStatusById(id)

	if err != nil {
		lib.HandlerNotfound(c, "Data not found")
		return
	}

	if transactionStatus.Id == 0 {
		lib.HandlerBadReq(c, "Failed to request")
		return
	}

	lib.HandlerOK(c, "Detail Order Types", transactionStatus, nil)
}

func GetALLTransactionStatus(c *gin.Context) {
	status, err := repository.FindAllTransactionStatus()

	if err != nil {
		lib.HandlerBadReq(c, "Failed to request")
		return
	}

	lib.HandlerOK(c, "List All Category", status, nil)
}

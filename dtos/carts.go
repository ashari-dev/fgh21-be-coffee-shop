package dtos

type FormCarts struct {
	TransactionDetail int `json:"transactionDetail" form:"transactionDetail"`
	Quantity          int `json:"quantity" form:"quantity" binding:"required"`
	Product           int `json:"product"`
	Variant           int `json:"variant" form:"variant" binding:"required"`
	ProductSize       int `json:"productSize" form:"productSize" binding:"required"`
}

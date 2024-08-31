package dtos

type FormTransaction struct{
	FullName string `json:"fullName" form:"fullName"`
	Email string `json:"email" form:"email"`
	Address string `json:"address" form:"address"`
	Payment string `json:"payment" form:"payment"`
	OrderType int `json:"orderType" form:"orderType"`
	TransactionStatus int `json:"transactionStatus" form:"transactionStatus"`
}
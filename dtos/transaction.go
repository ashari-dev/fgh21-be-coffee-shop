package dtos

type TransactionDetail struct {
	Quantity    int `json:"quantity" form:"quantity"`
	Product     int `json:"product"`
	Variant     int `json:"variant" form:"variant"`
	ProductSize int `json:"productSize" form:"productSize"`
}

type FormTransaction struct {
	FullName          string `json:"fullName" form:"fullName"`
	Email             string `json:"email" form:"email"`
	Address           string `json:"address" form:"address"`
	Payment           string `json:"payment" form:"payment"`
	TransactionDetail int    `json:"transactionDetail" form:"transactionDetail"`
	OrderType         int    `json:"orderType" form:"orderType"`
	TransactionStatus int    `json:"transactionStatus" form:"transactionStatus"`
}

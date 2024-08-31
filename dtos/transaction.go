package dtos


type TransactionDetail struct {
	Id            int `json:"id"`
	Quantity      int `json:"quantity"`
	ProductId     int `json:"productId"`
	TransactionId int `json:"transactionId"`
	VariantId     int `json:"variantId"`
	ProductSizeId int `json:"productSizeId"`
}

type FormTransaction struct{
	FullName string `json:"fullName" form:"fullName"`
	Email string `json:"email" form:"email"`
	Address string `json:"address" form:"address"`
	Payment string `json:"payment" form:"payment"`
	OrderType int `json:"orderType" form:"orderType"`
	TransactionStatus int `json:"transactionStatus" form:"transactionStatus"`
}


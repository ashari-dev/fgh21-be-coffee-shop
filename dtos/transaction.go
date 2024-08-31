package dtos

type TransactionDetail struct {
	Id            int `json:"id"`
	Quantity      int `json:"quantity"`
	ProductId     int `json:"productId"`
	TransactionId int `json:"transactionId"`
	VariantId     int `json:"variantId"`
	ProductSizeId int `json:"productSizeId"`
}

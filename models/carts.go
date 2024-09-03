package models

type Carts struct {
	Id                int `json:"id"`
	TransactionDetail int `json:"transactionDetail" db:"transaction_detail_id"`
	Quantity          int `json:"quantity" db:"quantity"`
	VariantId         int `json:"variantId" db:"variant_id"`
	ProductSizeId     int `json:"productSizeId" db:"sizes_id"`
	ProductId         int `json:"productId" db:"product_id"`
	UserId            int `json:"userId" db:"user_id"`
}
type CartsJoin struct {
	Id                int    `json:"id"`
	TransactionDetail int    `json:"transactionDetail"`
	Quantity          int    `json:"quantity"`
	Variant           string `json:"variant"`
	Size              string `json:"size"`
	Title             string `json:"title"`
	Price             int    `json:"price"`
}

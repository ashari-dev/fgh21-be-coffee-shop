package models

type TransactionDetail struct {
	Id            int `json:"id"`
	Quantity      int `json:"quantity"`
	ProductId     int `json:"productId" db:"product_id"`
	VariantId     int `json:"variantId" db:"variant_id"`
	ProductSizeId int `json:"productSizeId" db:"product_size_id"`
}

type TransactionProduct struct {
	NoOrder   int    `json:"noOrder"`
	Title     string `json:"title"`
	Quantity  int    `json:"quantity"`
	Variant   string `json:"variant"`
	Size      string `json:"size"`
	OrderType string `json:"orderType"`
	Price     int    `json:"price"`
}

type TransactionDetailJoin struct {
	NoOrder           int     `json:"noOrder"`
	FullName          string  `json:"fullName"`
	Address           string  `json:"address"`
	Payment           string  `json:"payment"`
	TransactionStatus string  `json:"transactionStatus"`
	Quantity          []int   `json:"quantity"`
	Price             []int   `json:"price"`
	OrderType         string  `json:"orderType"`
	PhoneNumber       *string `json:"phoneNumber"`
	Image string `json:"image"`
}

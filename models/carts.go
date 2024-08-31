package models

type Carts struct {
	Id           int `json:"id"`
	Quantity     int `json:"quantity" db:"quantity"`
	VariantId    int `json:"variantId" db:"variant_id"`
	SizesProduct int `json:"sizeProduct" db:"sizes_id"`
	ProductId    int `json:"productId" db:"product_id"`
	UserId       int `json:"userId" db:"user_id"`
}

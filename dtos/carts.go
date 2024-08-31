package dtos

type FormCarts struct {
	Quantity       int `json:"quantity" form:"quantity" binding:"required"`
	VariantProduct int `json:"variantProduct" form:"variantProduct" binding:"required"`
	SizesProduct   int `json:"sizesProduct" form:"sizesProduct" binding:"required"`
	ProductId      int `json:"productId" form:"productId" binding:"required"`
}

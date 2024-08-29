package dtos

type Products struct {
	Title 		string `form:"title" binding:"required"`
	Description string `form:"description" binding:"required"`
	Price 		int		`form:"price" binding:"required"`
}
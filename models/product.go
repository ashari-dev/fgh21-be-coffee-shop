package models

import (
	"RGT/konis/lib"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Product struct{
	Id int `json:"id"`
	Title string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	Price int `json:"price" db:"price"`
	UserId int `json:"user_id" db:"user_id"`
}
func FindAllProduct() []Product{
	db := lib.DB()
	defer db.Close(context.Background())

	rows, _ := db.Query(context.Background(),
		`select * from "products" order by "id" asc`,
	)
	product, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Product])
	if err != nil {
		fmt.Println(err)
	}

	return product
}
func FindOneProduct(id int) []Product {
	db := lib.DB()
	defer db.Close(context.Background())

	rows, _ := db.Query(context.Background(),
		`select * from "products" where "id" = $1`,id,
	)
	product, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Product])
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(product)
	return product
}
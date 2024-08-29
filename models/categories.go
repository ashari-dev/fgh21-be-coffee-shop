package models

import (
	"RGT/konis/lib"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Categories struct {
	Id   int    `json:"id"`
	Name string `json:"name" db:"name"`
}

func FindAllCategories() []Categories {
	db := lib.DB()
	defer db.Close(context.Background())

	rows, _ := db.Query(context.Background(),
		`select * from "categories"`,
	)
	category, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Categories])
	if err != nil {
		fmt.Println(err)
	}
	return category
	// select * from "categories" order by "id" asc
}

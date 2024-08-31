package repository

import (
	"RGT/konis/lib"
	"RGT/konis/models"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func FindAllCategories() []models.Categories {
	db := lib.DB()
	defer db.Close(context.Background())

	rows, _ := db.Query(context.Background(),
		`select * from "categories"`,
	)
	category, err := pgx.CollectRows(rows, pgx.RowToStructByPos[models.Categories])
	if err != nil {
		fmt.Println(err)
	}
	return category
	// select * from "categories" order by "id" asc
}

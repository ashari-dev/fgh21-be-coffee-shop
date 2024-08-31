package repository

import (
	"RGT/konis/lib"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type OrderTypes struct {
	Id       int    `json:"id"`
	Name     string `json:"name" db:"name"`
	AddPrice int    `json:"addPrice" db:"add_price"`
}

func FindAllOrderTypes() ([]OrderTypes, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	rows, _ := db.Query(context.Background(),
		`SELECT * FROM "order_types"`,
	)
	order, err := pgx.CollectRows(rows, pgx.RowToStructByPos[OrderTypes])
	if err != nil {
		fmt.Println(err)
	}
	return order, nil
}

func FindOneOrderTypes(id int) (OrderTypes, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `SELECT * FROM "order_types" where "id"=$1;`

	query, err := db.Query(context.Background(), sql, id)

	if err != nil {
		return OrderTypes{}, err
	}

	order, err := pgx.CollectOneRow(query, pgx.RowToStructByPos[OrderTypes])

	if err != nil {
		return OrderTypes{}, err
	}

	return order, err

}

package repository

import (
	"RGT/konis/lib"
	"RGT/konis/models"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func FindAllRoles() []models.Roles {
	db := lib.DB()
	defer db.Close(context.Background())

	rows, _ := db.Query(context.Background(),
		`select * from "roles" order by "id" asc`,
	)
	roles, err := pgx.CollectRows(rows, pgx.RowToStructByPos[models.Roles])
	if err != nil {
		fmt.Println(err)
	}

	return roles
}

func FindOneRoles(id int) models.Roles {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `select * from "roles" where id = $1`
	rows, _ := db.Query(context.Background(), sql, id)
	roles, err := pgx.CollectOneRow(rows, pgx.RowToStructByPos[models.Roles])
	if err != nil {
		fmt.Println(err)
	}

	return roles
}

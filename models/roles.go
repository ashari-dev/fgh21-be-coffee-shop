package models

import (
	"RGT/konis/lib"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Roles struct{
	Id            int`json:"id"`
	Name          string `json:"name" db:"name"`
}

func FindAllRoles() []Roles {
	db := lib.DB()
	defer db.Close(context.Background())

	rows, _ := db.Query(context.Background(),
		`select * from "roles" order by "id" asc`,
	)
	roles, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Roles])
	if err != nil {
		fmt.Println(err)
	}

	return roles
}
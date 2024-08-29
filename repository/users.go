package repository

import (
	"RGT/konis/lib"
	"RGT/konis/models"
	"context"

	"github.com/jackc/pgx/v5"
)

func FindAllUsers() ([]models.Users, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `SELECT * FROM users`

	rows, err := db.Query(context.Background(), sql)

	if err != nil {
		return []models.Users{}, err
	}

	users, err := pgx.CollectRows(rows, pgx.RowToStructByPos[models.Users])

	if err != nil {
		return []models.Users{}, err
	}

	return users, nil
}

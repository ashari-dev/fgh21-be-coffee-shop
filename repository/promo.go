package repository

import (
	"RGT/konis/lib"
	"RGT/konis/models"
	"context"

	"github.com/jackc/pgx/v5"
)

func GetAllPromo(data models.Promo) ([]models.Promo, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `SELECT * FROM promo`

	rows, err := db.Query(context.Background(), sql)

	if err != nil {
		return []models.Promo{}, err
	}

	dataPromo, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Promo])

	if err != nil {
		return []models.Promo{}, err
	}

	return dataPromo, err
}
func GetPromoById(id int) (models.Promo, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `SELECT * from promo WHERE id=$1`

	query, err := db.Query(context.Background(), sql, id)

	if err != nil {
		return models.Promo{}, err
	}

	dataPromo, err := pgx.CollectOneRow(query, pgx.RowToStructByName[models.Promo])

	if err != nil {
		return models.Promo{}, err
	}

	return dataPromo, err
}
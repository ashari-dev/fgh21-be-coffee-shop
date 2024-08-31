package repository

import (
	"RGT/konis/lib"
	"RGT/konis/models"
	"context"

	"github.com/jackc/pgx/v5"
)

func GetAllProductsSize(data models.ProductsSizes) ([]models.ProductsSizes, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `SELECT * FROM product_sizes`

	rows, err := db.Query(context.Background(), sql)

	if err != nil {
		return []models.ProductsSizes{}, err
	}

	productsSizes, err := pgx.CollectRows(rows, pgx.RowToStructByPos[models.ProductsSizes])

	if err != nil {
		return []models.ProductsSizes{}, err
	}

	return productsSizes, err
}

func FindProductSizeByProductId(product_id int) ([]models.ProductsSizes, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `SELECT * FROM product_sizes WHERE product_id=$1`

	row, err := db.Query(context.Background(), sql, product_id)

	if err != nil {
		return []models.ProductsSizes{}, err
	}

	product_sizes, err := pgx.CollectRows(row, pgx.RowToStructByPos[models.ProductsSizes])

	if err != nil {
		return []models.ProductsSizes{}, err
	}

	return product_sizes, nil
}

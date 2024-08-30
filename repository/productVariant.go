package repository

import (
	"RGT/konis/lib"
	"RGT/konis/models"
	"context"

	"github.com/jackc/pgx/v5"
)
func GetAllProductVariant(data models.ProductVariant) ([]models.ProductVariant, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `SELECT * FROM product_variants`

	rows, err := db.Query(context.Background(), sql)

	if err != nil {
		return []models.ProductVariant{}, err
	}

	product_variant, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.ProductVariant])

	if err != nil {
		return []models.ProductVariant{}, err
	}

	return product_variant, err
}
func GetProductVariantById(product_id int) (models.ProductVariant, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `SELECT * FROM product_variants WHERE product_id=$1`

	row, err := db.Query(context.Background(), sql, product_id)

	if err != nil {
		return models.ProductVariant{}, err
	}

	product_variant, err := pgx.CollectOneRow(row, pgx.RowToStructByName[models.ProductVariant])

	if err != nil {
		return models.ProductVariant{}, err
	}

	return product_variant, nil
}

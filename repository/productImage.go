package repository

import (
	"RGT/konis/lib"
	"RGT/konis/models"
	"context"

	"github.com/jackc/pgx/v5"
)

func UploadProductImage(data models.ProductImage) (models.ProductImage, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `INSERT INTO product_images (image, product_id) VALUES ($1, $2) RETURNING *`

	row, err := db.Query(context.Background(), sql, data.Image, data.ProductId)
	if err != nil {
		return models.ProductImage{}, err
	}

	productImage, err := pgx.CollectOneRow(row, pgx.RowToStructByName[models.ProductImage])
	if err != nil {
		return models.ProductImage{}, err
	}

	return productImage, nil
}

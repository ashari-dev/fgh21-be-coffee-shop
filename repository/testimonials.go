package repository

import (
	"RGT/konis/lib"
	"RGT/konis/models"
	"context"

	"github.com/jackc/pgx/v5"
)

func FindAllTestimonials(page int) (models.Testimonials, error) {
	db := lib.DB()
	defer db.Close(context.Background())
	var offset int = (page - 1) * 1

	sql := `SELECT * FROM testimonials limit 1 offset ($1)`

	rows, err := db.Query(context.Background(), sql, offset)

	if err != nil {
		return models.Testimonials{}, err
	}

	testimonials, err := pgx.CollectOneRow(rows, pgx.RowToStructByPos[models.Testimonials])

	if err != nil {
		return models.Testimonials{}, err
	}

	return testimonials, nil
}

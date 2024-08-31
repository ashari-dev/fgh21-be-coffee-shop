package repository

import (
	"RGT/konis/lib"
	"RGT/konis/models"
	"context"

	"github.com/jackc/pgx/v5"
)

func FindAllTestimonials() ([]models.Testimonials, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `SELECT * FROM testimonials`

	rows, err := db.Query(context.Background(), sql)

	if err != nil {
		return []models.Testimonials{}, err
	}

	testimonials, err := pgx.CollectRows(rows, pgx.RowToStructByPos[models.Testimonials])

	if err != nil {
		return []models.Testimonials{}, err
	}

	return testimonials, nil
}

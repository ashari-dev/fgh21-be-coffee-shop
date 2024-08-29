package repository

import (
	"RGT/konis/lib"
	"RGT/konis/models"
	"context"

	"github.com/jackc/pgx/v5"
)

func FindProfileById(id int) (models.ProfileJoinUser, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `
		SELECT p.id, p.full_name, u.email, p.phone_number,
		p.address, p.image, u.role_id 
		FROM profile p 
		JOIN users u ON u.id = p.user_id
		WHERE p.id = $1
		`

	row, err := db.Query(context.Background(), sql, id)

	if err != nil {
		return models.ProfileJoinUser{}, err
	}

	data, err := pgx.CollectOneRow(row, pgx.RowToStructByPos[models.ProfileJoinUser])

	if err != nil {
		return models.ProfileJoinUser{}, err
	}

	return data, nil
}

func CreateProfile(data models.Profile) (models.Profile, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `
		INSERT INTO profile (full_name, user_id)
		VALUES ($1, $2) RETURNING *
		`
	row, err := db.Query(context.Background(), sql, data.FullName, data.UserId)

	if err != nil {
		return models.Profile{}, nil
	}

	profile, err := pgx.CollectOneRow(row, pgx.RowToStructByPos[models.Profile])

	if err != nil {
		return models.Profile{}, nil
	}

	return profile, err
}

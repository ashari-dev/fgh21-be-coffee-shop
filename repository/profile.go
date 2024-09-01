package repository

import (
	"RGT/konis/dtos"
	"RGT/konis/lib"
	"RGT/konis/models"
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

func FindProfileById(id int) (dtos.ProfileUser, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `
		SELECT p.id, p.full_name, u.email, p.phone_number,
		p.address, p.image, u.role_id 
		FROM profile p 
		JOIN users u ON u.id = p.user_id
		WHERE p.user_id = $1
		`

	row, err := db.Query(context.Background(), sql,id)

	if err != nil {
		return dtos.ProfileUser{}, err
	}

	data, err := pgx.CollectOneRow(row, pgx.RowToStructByName[dtos.ProfileUser])

	if err != nil {
		log.Println(err)
		fmt.Println(err)
		return dtos.ProfileUser{}, err
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

func UpdateProfile(data models.Profile, id int) (dtos.ProfileJoinUser, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `UPDATE profile SET ("full_name", "phone_number", "address") = ($1, $2, $3) WHERE user_id=$4 returning "id", "full_name", "phone_number", "address"`

	query := db.QueryRow(context.Background(), sql, data.FullName, data.PhoneNumber, data.Address, id)

	var result dtos.ProfileJoinUser
	err := query.Scan(
		&result.Id,
		&result.FullName,
		&result.PhoneNumber,
		&result.Address,
		// &result.Image,
	)
	fmt.Println(err)

	if err != nil {
		return dtos.ProfileJoinUser{}, err
	}

	return result, err
}

func RemoveProfile(id int) error {
	db := lib.DB()
	defer db.Close(context.Background())

	// profileDelete, err := FindProfileById(id)
	// if err != nil {
	// 	return dtos.ProfileJoinUser{}, err
	// }

	sql := `DELETE FROM profile WHERE id=$1;`

	db.Exec(context.Background(), sql, id)

	// if query.RowsAffected() == 0 {
	// 	return fmt.Errorf("data not found")
	// }

	return nil
}

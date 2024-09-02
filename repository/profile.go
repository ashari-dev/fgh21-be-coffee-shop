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

// Id       int    `json:"id"`
// FullName string `json:"fullName" db:"full_name"`
// Email    string `json:"email" `
// PhoneNumber string  `json:"phoneNumber" db:"phone_number"`
// Address     string  `json:"address"`
// Image       *string `json:"image"`

func FindAllProfiles(search string, page int, limit int) ([]models.ProfileJoinUser, int) {
	db := lib.DB()
	defer db.Close(context.Background())
	offset := 0
	if page > 1 {
		offset = (page - 1) * limit
	}

	// sql := `SELECT p.id, p.full_name, p.phone_number, u.email,
	// 	p.address, p.image
	// 	FROM profile p
	// 	JOIN users u ON u.id = p.user_id`

	sql := `SELECT p.id, p.full_name, p.phone_number, u.email, p.address, p.image 
	FROM profile p
	JOIN users u ON u.id = p.user_id
	where "full_name" ilike '%' || $1 || '%'
	limit $2 offset $3`

	rows, _ := db.Query(context.Background(),
		sql, search, limit, offset,
	)
	count := TotalProfile(search)

	profile, err := pgx.CollectRows(rows, pgx.RowToStructByPos[models.ProfileJoinUser])
	if err != nil {
		return []models.ProfileJoinUser{}, count
	}

	return profile, count
}

func TotalProfile(search string) int {
	db := lib.DB()
	defer db.Close(context.Background())
	inputSQL := `select count(id) as "total" from "profile" where "full_name" ilike '%' || $1 || '%'`
	rows := db.QueryRow(context.Background(), inputSQL, search)
	var result int
	rows.Scan(
		&result,
	)
	return result
}

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

	row, err := db.Query(context.Background(), sql, id)

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

// Id          int     `json:"id"`
// 	FullName    string  `json:"fullName" db:"full_name"`
// 	PhoneNumber *string `json:"phoneNumber" db:"phone_number"`
// 	Address     *string `json:"address"`
// 	Image       *string `json:"image"`
// 	UserId      int     `json:"userId" db:"user_id"`

func CreateProfileJoinUser(data models.Profile) (models.Profile, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `
		INSERT INTO profile (full_name, phone_number, address, image, user_id)
		VALUES ($1, $2, $3, $4, $5) RETURNING *
		`
	row, err := db.Query(context.Background(), sql, data.FullName, data.PhoneNumber, data.Address, data.Image, data.UserId)

	if err != nil {
		fmt.Println(err)
		return models.Profile{}, nil
	}

	profile, err := pgx.CollectOneRow(row, pgx.RowToStructByPos[models.Profile])

	if err != nil {
		// fmt.Println(err)
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

func UpdateProfileImage(data models.Profile, id int) (models.Profile, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `UPDATE profile SET "image" = $1 WHERE user_id=$2 returning *`

	row, err := db.Query(context.Background(), sql, data.Image, id)
	if err != nil {
		return models.Profile{}, nil
	}

	profile, err := pgx.CollectOneRow(row, pgx.RowToStructByName[models.Profile])
	if err != nil {
		return models.Profile{}, nil
	}

	return profile, nil
}

func RemoveProfile(id int) error {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `DELETE FROM profile WHERE id=$1;`

	db.Exec(context.Background(), sql, id)

	return nil
}

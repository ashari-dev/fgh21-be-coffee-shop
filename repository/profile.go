package repository

import (
	"RGT/konis/dtos"
	"RGT/konis/lib"
	"RGT/konis/models"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

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

	sql := `
		UPDATE profile 
		SET ("full_name", "phone_number", "address") = ($1, $2, $3)
		WHERE user_id=$4
		RETURNING
		"id", "full_name","phone_number",
		"address", "image"`

	query := db.QueryRow(context.Background(), sql, data.FullName, data.PhoneNumber, data.Address, id)

	var result dtos.ProfileJoinUser
	err := query.Scan(
		&result.Id,
		&result.FullName,
		&result.PhoneNumber,
		&result.Address,
		&result.Image,
		// &result.RoleId,
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

func DeleteProfileAndUser(id int) (models.Users, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	tx, err := db.Begin(context.Background())
	if err != nil {
		return models.Users{}, fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback(context.Background())
		} else {
			tx.Commit(context.Background())
		}
	}()

	var userId int
	var profileImage *string
	err = tx.QueryRow(context.Background(), `
		DELETE FROM profile 
		WHERE id = $1 
		RETURNING user_id, image`, id).Scan(&userId, &profileImage)
	if err != nil {
		if err == pgx.ErrNoRows {
			return models.Users{}, fmt.Errorf("profile not found")
		}
		return models.Users{}, fmt.Errorf("failed to delete profile or find associated user: %v", err)
	}

	var user models.Users
	err = tx.QueryRow(context.Background(), `
		DELETE FROM users 
		WHERE id = $1 
		RETURNING id, email, role_id`, userId).Scan(&user.Id, &user.Email, &user.RoleId)
	if err != nil {
		return models.Users{}, fmt.Errorf("failed to delete user: %v", err)
	}

	if profileImage != nil {

		filePathParts := strings.Split(*profileImage, "8000")
		if len(filePathParts) > 1 {
			filePath := "." + filePathParts[1]
			err = os.Remove(filePath)
			if err != nil && !os.IsNotExist(err) {

				fmt.Printf("Warning: failed to remove image file: %v\n", err)
			}
		}
	}

	return user, nil
}

package repository

import (
	"RGT/konis/lib"
	"RGT/konis/models"
	"context"

	"github.com/jackc/pgx/v5"
)

func FindAllUsers() ([]models.Users, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `SELECT * FROM users`

	rows, err := db.Query(context.Background(), sql)

	if err != nil {
		return []models.Users{}, err
	}

	users, err := pgx.CollectRows(rows, pgx.RowToStructByPos[models.Users])

	if err != nil {
		return []models.Users{}, err
	}

	return users, nil
}

func FindUserById(id int) (models.Users, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `SELECT * FROM users WHERE id=$1`

	row, err := db.Query(context.Background(), sql, id)

	if err != nil {
		return models.Users{}, err
	}

	user, err := pgx.CollectOneRow(row, pgx.RowToStructByPos[models.Users])

	if err != nil {
		return models.Users{}, err
	}

	return user, nil
}

func FindUserByEmail(email string) (models.Users, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `SELECT * FROM users WHERE email=$1`

	row, err := db.Query(context.Background(), sql, email)

	if err != nil {
		return models.Users{}, err
	}

	user, err := pgx.CollectOneRow(row, pgx.RowToStructByPos[models.Users])

	if err != nil {
		return models.Users{}, err
	}

	return user, nil
}

func CreateUser(data models.Users) (models.Users, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	data.Password = lib.Encrypt(data.Password)

	sql := `
		INSERT INTO users (email, password, role_id)
		VALUES ($1, $2, $3) RETURNING *
		`
	row, err := db.Query(context.Background(), sql, data.Email, data.Password, data.RoleId)

	if err != nil {
		return models.Users{}, nil
	}

	user, err := pgx.CollectOneRow(row, pgx.RowToStructByPos[models.Users])
	// fmt.Println(user)
	if err != nil {
		return models.Users{}, nil
	}

	return user, err
}

func UpdateUserById(data models.Users, id int) (models.Users, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	if data.Password != "" {
		data.Password = lib.Encrypt(data.Password)
	}

	sql := `
		UPDATE users SET(email, password)=(COALESCE(NULLIF($1,''),"email"), COALESCE(NULLIF($2,''),"password")) 
		WHERE id = $3 RETURNING *
		`

	row, err := db.Query(context.Background(), sql, data.Email, data.Password, id)

	if err != nil {
		return models.Users{}, nil
	}

	user, err := pgx.CollectOneRow(row, pgx.RowToStructByPos[models.Users])
	if err != nil {
		return models.Users{}, nil
	}

	return user, err
}

func DeleteUserById(id int) (models.Users, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	userDelete, err := FindUserById(id)
	if err != nil {
		return models.Users{}, err
	}
	sql := `DELETE FROM users WHERE id=$1`

	db.Exec(context.Background(), sql, id)

	return userDelete, nil
}

func CreateinsertUser(user models.InsertUsers) (models.InsertUsers, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `INSERT INTO users (email, password, role_id) VALUES ($1, $2, $3) RETURNING id, email, role_id`
	row := db.QueryRow(context.Background(), sql, user.Email, user.Password, user.RoleId)

	var createdUser models.InsertUsers
	err := row.Scan(&createdUser.Id, &createdUser.Email, &createdUser.RoleId)
	if err != nil {
		return models.InsertUsers{}, err
	}

	return createdUser, nil
}

func CreateinsertProfile(profile models.InsertProfile) (models.InsertProfile, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `
		INSERT INTO profile (full_name, phone_number, address, image, user_id)
		VALUES ($1, $2, $3, $4, $5) 
		RETURNING id, full_name, phone_number, address, image, user_id
	`
	row := db.QueryRow(context.Background(), sql, profile.FullName, profile.PhoneNumber, profile.Address, profile.Image, profile.UserId)

	var createdProfile models.InsertProfile
	err := row.Scan(&createdProfile.Id, &createdProfile.FullName, &createdProfile.PhoneNumber, &createdProfile.Address, &createdProfile.Image, &createdProfile.UserId)
	if err != nil {
		return models.InsertProfile{}, err
	}

	return createdProfile, nil
}

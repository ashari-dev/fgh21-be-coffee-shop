package repository

import (
	"RGT/konis/dtos"
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

func CreateUser(user dtos.User, roleId int) models.Users {
	db := lib.DB()
	defer db.Close(context.Background())
	// user.Password = lib.Encrypt(user.Password)
	// fmt.Println(user)

	row := db.QueryRow(
		context.Background(),
		`insert into "users" (email, password, role_id) values ($1, $2, $3) returning "id", "email", "password", "role_id"`,
		user.Email, user.Password, roleId,
	)

	var results models.Users
	row.Scan(
		&results.Id,
		&results.Email,
		&results.Password,
		&results.RoleId,
	)
	return results
}

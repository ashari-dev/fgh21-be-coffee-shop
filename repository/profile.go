package repository

import (
	"RGT/konis/dtos"
	"RGT/konis/lib"
	"RGT/konis/models"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func CreateProfile(data models.JoinProfile, roleId int) dtos.Profile {
	db := lib.DB()
	defer db.Close(context.Background())

	var profile dtos.Profile

	// data.Password = lib.Encrypt(data.Password)

	sqlRegist := `insert into "users" 
	("email", "password", "role_id") 
	values 
	($1, $2, $3) returning "id", "email", "role_id"`

	// var userId int
	row1 := db.QueryRow(context.Background(), sqlRegist, data.Email, data.Password, roleId)
	row1.Scan(
		&data.Id,
		&data.Email,
		&data.RoleId,
	)
	// fmt.Println(err1)
	// if err1 != nil {
	// 	fmt.Println(err1)
	// 	fmt.Println("err1")
	// }

	sqlProfile := `insert into "profile" 
	("full_name","phone_number", "address", "image", "user_id") 
	values 
	($1, $2, $3, $4, $5) returning "id", "full_name", "phone_number", "address", "image", "user_id"`

	row2 := db.QueryRow(context.Background(), sqlProfile, data.Profile.FullName, data.Profile.PhoneNumber, data.Profile.Address, data.Profile.Image, data.Profile.UserId)
	row2.Scan(
		&profile.Id,
		&profile.FullName,
		&profile.PhoneNumber,
		&profile.Address,
		&profile.Image,
		&profile.UserId,
	)

	// if err2 != nil {
	// 	fmt.Println(err2)
	// }
	return profile
}

func FindAllProfiles() []dtos.Profile {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `select * from "profile"`

	rows, err := db.Query(context.Background(), sql)

	if err != nil {
		fmt.Println(err)
	}

	profiles, err := pgx.CollectRows(rows, pgx.RowToStructByPos[dtos.Profile])
	if err != nil {
		fmt.Println(err)
	}

	return profiles
}

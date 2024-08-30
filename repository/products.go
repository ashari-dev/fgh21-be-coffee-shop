package repository

import (
	"RGT/konis/lib"
	"RGT/konis/models"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func GetAllProducts(data models.Products) ([]models.Products, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `SELECT * FROM products`

	rows, err := db.Query(context.Background(), sql)

	if err != nil {
		return []models.Products{}, err
	}

	products, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Products])

	if err != nil {
		return []models.Products{}, err
	}

	return products, err
}
func AddNewProduct(data models.Products) (models.Products, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `INSERT into products ("title", "description", "price") VALUES ($1, $2, $3) returning id, "title", "description", "price", "user_id"`

	query, err := db.Query(context.Background(), sql, data.Title, data.Description, data.Price)

	if err != nil {
		return models.Products{}, err
	}

	row, err := pgx.CollectOneRow(query, pgx.RowToStructByName[models.Products])

	if err != nil {
		fmt.Println(err)
		return models.Products{}, err
	}

	return row, err
}
func GetProductById(id int) (models.Products, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `SELECT * from products WHERE id=$1;`

	query, err := db.Query(context.Background(), sql, id)

	if err != nil {
		return models.Products{}, err
	}

	selectedRow, err := pgx.CollectOneRow(query, pgx.RowToStructByName[models.Products])

	if err != nil {
		return models.Products{}, err
	}

	return selectedRow, err
}
func ChangeDataProduct(data models.Products, id int) (models.Products, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `UPDATE products SET ("title", "description", "price") = ($1, $2, $3) WHERE id=$4 returning "id", "title", "description", "price"`

	query := db.QueryRow(context.Background(), sql, data.Title, data.Description, data.Price, id)

	var result models.Products

	err := query.Scan(
		&result.Id,
		&result.Title,
		&result.Description,
		&result.Price,
	)

	if err != nil {
		return models.Products{}, err
	}

	return result, err
}
func RemoveTheProduct(data models.Products, id int) error {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `DELETE FROM products WHERE id=$1;`

	query, _ := db.Exec(context.Background(), sql, id)

	// if err != nil {
	// 	// return fmt.Errorf("Failed to delete product")
	// }
	if query.RowsAffected() == 0 {
		return fmt.Errorf("data not found")
	}

	return nil
}

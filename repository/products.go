package repository

import (
	"RGT/konis/lib"
	"RGT/konis/models"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func GetAllProducts(page int, limit int) ([]models.JoinProducts, error) {
	db := lib.DB()
	defer db.Close(context.Background())
	var offset int = (page - 1) * limit

	sql := `SELECT p.id, pi.image, p.title, p.price, p.description, array_agg(ps.id) as "product_sizes", array_agg(pt.order_type_id) as "order_type", pv.stock
	FROM products p
	JOIN product_images pi ON pi.product_id = p.id
	JOIN product_sizes ps ON ps.product_id = p.id
	JOIN product_order_types pt ON pt.product_id = p.id
	JOIN product_variants pv ON pv.product_id = p.id
	GROUP BY p.id, pi.image, p.title, p.description, pv.stock
	limit $1 offset $2
	`

	// sql := `SELECT * FROM products limit $1 offset $2`

	rows, err := db.Query(context.Background(), sql, limit, offset)

	if err != nil {
		return []models.JoinProducts{}, err
	}

	products, err := pgx.CollectRows(rows, pgx.RowToStructByPos[models.JoinProducts])

	if err != nil {
		return []models.JoinProducts{}, err
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

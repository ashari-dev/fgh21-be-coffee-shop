package repository

import (
	"RGT/konis/lib"
	"RGT/konis/models"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)


func GetAllProducts(page int, limit int) ([]models.Product, error) {
	db := lib.DB()
	defer db.Close(context.Background())
	// var offset int = (page - 1) * limit



	// 	sql := `SELECT p.id, pi.image, p.title, p.price, p.description, array_agg(ps.id) as "product_sizes", array_agg(pt.order_type_id) as "order_type", pv.stock
	// 	FROM products p
	// 	JOIN product_images pi ON pi.product_id = p.id
	// 	JOIN product_sizes ps ON ps.product_id = p.id
	// 	JOIN product_order_types pt ON pt.product_id = p.id
	// 	JOIN product_variants pv ON pv.product_id = p.id
	// 	GROUP BY p.id, pi.image, p.title, p.description, pv.stock
	// 	`


	sql := `SELECT * FROM products limit $1 offset $2`

	rows, err := db.Query(context.Background(), sql)

	if err != nil {
		return []models.Product{}, err
	}

	products, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Product])

	if err != nil {
		return []models.Product{}, err

	}

	return products, err
}
func GetAllProductsWithPagination(page int, limit int) ([]models.Products, error) {
	db := lib.DB()
	defer db.Close(context.Background())
	var offset int = (page - 1) * limit

	sql := `SELECT * FROM products limit $1 offset $2`

	rows, err := db.Query(context.Background(), sql, limit, offset)

	if err != nil {
		return []models.Products{}, err
	}

	products, err := pgx.CollectRows(rows, pgx.RowToStructByPos[models.Products])

	if err != nil {
		return []models.Products{}, err
	}

	return products, err
}

func AddNewProduct(data models.Products) (models.Products, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `INSERT into products ("title", "description", "price", "stock", "user_id") VALUES ($1, $2, $3, $4, $5) returning id, "title", "description", "price", "stock", "user_id"`

	query, err := db.Query(context.Background(), sql, data.Title, data.Description, data.Price, data.Stock, data.UserId)

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

	sql := `UPDATE products SET ("title", "description", "price", "stock") = ($1, $2, $3, $4) WHERE id=$5 returning "id", "title", "description", "price", "stock"`

	query := db.QueryRow(context.Background(), sql, data.Title, data.Description, data.Price, data.Stock, id)

	var result models.Products

	err := query.Scan(
		&result.Id,
		&result.Title,
		&result.Description,
		&result.Price,
		&result.Stock,
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
func GetAllProductsWithPagination(page int, limit int) ([]models.Products, error) {
	db := lib.DB()
	defer db.Close(context.Background())
	var offset int = (page - 1) * limit

	sql := `SELECT * FROM products limit $1 offset $2`

	rows, err := db.Query(context.Background(), sql, limit, offset)

	if err != nil {
		return []models.Products{}, err
	}

	products, err := pgx.CollectRows(rows, pgx.RowToStructByPos[models.Products])

	if err != nil {
		return []models.Products{}, err
	}

	return products, err
}

// func FilterProduct(dt) {
// 	db := lib.DB()
// 	defer db.Close(context.Background())

// 	sql := `select * from "products" where "price" >= $1 and "price" <= $2;`

// 	query, err := db.Query(context.Background(), sql, lowPrice, highPrice,)

// 	if err != nil {
// 		return models.Products{}, err
// 	}

// 	selectedRow, err := pgx.CollectOneRow(query, pgx.RowToStructByName[models.Products])

// 	if err != nil {
// 		return models.Products{}, err
// 	}

// 	return selectedRow, err
// }

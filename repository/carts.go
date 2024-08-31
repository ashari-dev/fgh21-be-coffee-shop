package repository

import (
	"RGT/konis/lib"
	"RGT/konis/models"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func FindAllCarts(id int) ([]models.Carts, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	// sql := `
	// 	select c.id, c.quantity, pv.id, p.id, u.id from carts c
	// 	join product_variants pv on pv.id = c.variant_id
	// 	join products p on p.id = c.product_id
	// 	join users u on u.id = c.user_id
	// 	where c.id = 1;`

	sql := `SELECT * from carts WHERE "user_id" = $1`

	query, err := db.Query(context.Background(), sql, id)

	if err != nil {
		return []models.Carts{}, err
	}

	rows, err := pgx.CollectRows(query, pgx.RowToStructByPos[models.Carts])

	if err != nil {
		return []models.Carts{}, err
	}
	fmt.Println(rows)

	return rows, err
}
func CreateCarts(data models.Carts) (models.Carts, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	fmt.Println(data)
	// fmt.Println(userId)

	sql := `INSERT INTO carts ("quantity", "variant_id", "sizes_id", "product_id", "user_id") VALUES ($1, $2, $3, $4, $5) RETURNING "id", "quantity", "variant_id", "sizes_id", "product_id", "user_id"`

	// sql := `insert into carts "quantity" values $1 returning "id", "quantity", "variant_id", "sizes_id", "product_id", "user_id"`

	row := db.QueryRow(context.Background(), sql, data.Quantity, data.VariantId, data.SizesProduct, data.ProductId, data.UserId)
	// row, err := db.Query(context.Background(), sql, data.Quantity)

	// if err != nil {
	// 	return models.Carts{}, nil
	// }

	// results, err := pgx.CollectOneRow(row, pgx.RowToStructByPos[models.Carts])

	// if err != nil {
	// 	return models.Carts{}, nil
	// }

	var results models.Carts

	// Id           int `json:"id"`
	// Quantity     int `json:"quantity"`
	// VariantId    int `json:"variantId" db:"variant_id"`
	// SizesProduct int `json:"sizeProduct" db:"sizes_id"`
	// ProductId    int `json:"productId" db:"product_id"`
	// UserId       int `json:"userId" db:"user_id"`

	row.Scan(
		&results.Id,
		&results.Quantity,
		&results.VariantId,
		&results.SizesProduct,
		&results.ProductId,
		&results.UserId,
	)
	fmt.Println(results)
	return results, nil
}

func GetCartsById(id int) (models.Carts, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `SELECT * from carts WHERE id=$1`

	query, err := db.Query(context.Background(), sql, id)

	if err != nil {
		return models.Carts{}, err
	}

	selectedRow, err := pgx.CollectOneRow(query, pgx.RowToStructByName[models.Carts])

	if err != nil {
		return models.Carts{}, err
	}

	return selectedRow, err
}

func DeleteCarts(data models.Carts, id int) error {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `DELETE FROM carts WHERE id=$1`

	query, _ := db.Exec(context.Background(), sql, id)

	// if err != nil {
	// 	// return fmt.Errorf("Failed to delete product")
	// }
	if query.RowsAffected() == 0 {
		return fmt.Errorf("data not found")
	}

	return nil
}

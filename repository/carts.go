package repository

import (
	"RGT/konis/lib"
	"RGT/konis/models"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func FindAllCarts(id int) ([]models.CartsJoin, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	// sql := `
	// 	select c.id, c.quantity, pv.id, p.id, u.id from carts c
	// 	join product_variants pv on pv.id = c.variant_id
	// 	join products p on p.id = c.product_id
	// 	join users u on u.id = c.user_id
	// 	where c.id = 1;`

	sql := `SELECT carts.id, carts.transaction_detail_id, carts.quantity, product_variants.name as variant, product_sizes.name as size, products.title, products.price  FROM carts
			INNER JOIN product_variants ON carts.variant_id = product_variants.id
			INNER JOIN product_sizes ON carts.sizes_id = product_sizes.id
			INNER JOIN products on carts.product_id = products.id
			WHERE carts.user_id = $1`

	query, err := db.Query(context.Background(), sql, id)

	if err != nil {
		return []models.CartsJoin{}, err
	}

	rows, err := pgx.CollectRows(query, pgx.RowToStructByPos[models.CartsJoin])

	if err != nil {
		return []models.CartsJoin{}, err
	}
	fmt.Println(rows)

	return rows, err
}
func CreateCarts(data models.Carts) (models.Carts, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	fmt.Println(data)
	// fmt.Println(userId)

	sql := `INSERT INTO carts ("transaction_detail_id", "quantity", "variant_id", "sizes_id", "product_id", "user_id") VALUES ($1, $2, $3, $4, $5, $6) RETURNING "id", "transaction_detail_id", "quantity", "variant_id", "sizes_id", "product_id", "user_id"`

	// sql := `insert into carts "quantity" values $1 returning "id", "quantity", "variant_id", "sizes_id", "product_id", "user_id"`

	row, _ := db.Query(context.Background(), sql, data.TransactionDetail, data.Quantity, data.VariantId, data.ProductSizeId, data.ProductId, data.UserId)
	results, err := pgx.CollectOneRow(row, pgx.RowToStructByPos[models.Carts])
	if err != nil {
		fmt.Println(err)
	}
	// row, err := db.Query(context.Background(), sql, data.Quantity)

	// if err != nil {
	// 	return models.Carts{}, nil
	// }

	// if err != nil {
	// 	return models.Carts{}, nil
	// }

	// var results models.Carts

	// Id           int `json:"id"`
	// Quantity     int `json:"quantity"`
	// VariantId    int `json:"variantId" db:"variant_id"`
	// SizesProduct int `json:"sizeProduct" db:"sizes_id"`
	// ProductId    int `json:"productId" db:"product_id"`
	// UserId       int `json:"userId" db:"user_id"`

	// row.Scan(
	// 	&results.Id,
	// 	&results.TransactionDetail,
	// 	&results.Quantity,
	// 	&results.VariantId,
	// 	&results.ProductSizeId,
	// 	&results.ProductId,
	// 	&results.UserId,
	// )
	fmt.Println("ini results")
	fmt.Println(results)
	return results, nil
}

func GetCartsByUserId(id int) ([]models.Carts, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `SELECT * from carts WHERE user_id=$1`

	query, err := db.Query(context.Background(), sql, id)

	if err != nil {
		return []models.Carts{}, err
	}

	selectedRow, err := pgx.CollectRows(query, pgx.RowToStructByName[models.Carts])

	if err != nil {
		return []models.Carts{}, err
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

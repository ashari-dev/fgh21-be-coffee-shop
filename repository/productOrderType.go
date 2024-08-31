package repository

import (
	"RGT/konis/lib"
	"RGT/konis/models"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func FindAllProductOrderType (models.ProductOrderType) ([]models.ProductOrderType, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `SELECT * from product_order_types`

	query, err := db.Query(context.Background(), sql)

	if err != nil {
		return []models.ProductOrderType{}, err
	}

	rows, err := pgx.CollectRows(query, pgx.RowToStructByPos[models.ProductOrderType])

	if err != nil {
		return []models.ProductOrderType{}, err
	}

	return rows, err
}
func AddNewProductOrderType (data models.ProductOrderType) (models.ProductOrderType, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `INSERT into product_order_types ("product_id", "order_type_id") VALUES ($1, $2)`

	query, err := db.Query(context.Background(), sql, data.ProductId, data.OrderTypeId)

	if err != nil {
		return models.ProductOrderType{}, err
	}

	row, err := pgx.CollectOneRow(query, pgx.RowToStructByName[models.ProductOrderType])

	if err != nil {
		return models.ProductOrderType{}, err
	}

	return row, err
}
func FindProductOrderTypeById (id int) (models.ProductOrderType, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `SELECT * from product_order_types WHERE id=$1`

	query := db.QueryRow(context.Background(), sql, id)

	var result models.ProductOrderType

	err := query.Scan(
		&result.Id,
		&result.ProductId,
		&result.OrderTypeId,
	)

	if err != nil {
		return models.ProductOrderType{}, err
	}

	return result, err
}
func EditProductOrderType(data models.ProductOrderType, id int) (models.ProductOrderType, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `UPDATE product_order_types SET ("product_id", "order_type_id") = ($1, $2) WHERE id=$3 returning "id", "product_id", "order_type_id"`

	query := db.QueryRow(context.Background(), sql, data.ProductId, data.OrderTypeId, id)

	var result models.ProductOrderType

	err := query.Scan(
		&result.Id,
		&result.ProductId,
		&result.OrderTypeId,
	)

	if err != nil {
		return models.ProductOrderType{}, err
	}

	return result, err
}
func RemoveProductOrderType (data models.ProductOrderType, id int) error {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `DELETE FROM product_order_types WHERE id=$1;`

	query, _ := db.Exec(context.Background(), sql, id)

	// if err != nil {
	// 	// return fmt.Errorf("Failed to delete product")
	// }
	if query.RowsAffected() == 0 {
		return fmt.Errorf("data not found")
	}

	return nil
}
package repository

import (
	"RGT/konis/lib"
	"RGT/konis/models"
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

func CreateTransactionDetail(data models.TransactionDetail) (models.TransactionDetail, error) {
	db := lib.DB()
	defer db.Close(context.Background())
	fmt.Println(data)
	sql := `
		INSERT INTO transaction_details 
		(quantity, product_id, variant_id, product_size_id)
		VALUES
		($1, $2, $3, $4) RETURNING *
	`

	row, err := db.Query(context.Background(), sql, data.Quantity, data.ProductId, data.VariantId, data.ProductSizeId)
	if err != nil {
		return models.TransactionDetail{}, err
	}

	transactionDetail, err := pgx.CollectOneRow(row, pgx.RowToStructByPos[models.TransactionDetail])

	if err != nil {
		return models.TransactionDetail{}, err
	}

	return transactionDetail, err
}

func FindTransactionDetailById(id int) (models.TransactionDetailJoin, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `SELECT  transactions.no_order, transactions.add_full_name, transactions.add_address, transactions.payment , transaction_status.name AS transaction_status, transaction_details.quantity, products.price, order_types.name AS order_type, profile.phone_number ,product_images.image
	FROM transaction_details
	INNER JOIN transactions ON transactions.transaction_detail_id = transaction_details.id
	INNER JOIN transaction_status on transactions.transaction_status_id = transaction_status.id
	INNER JOIN order_types on transactions.order_type_id = order_types.id
	INNER JOIN profile on transactions.user_id = profile.user_id
	INNER JOIN products on transaction_details.product_id = products.id
	INNER JOIN product_images on product_images.product_id  = products.id
  WHERE no_order = $1
--     WHERE no_order = $1
--     GROUP BY transactions.no_order, transactions.add_full_name, transactions.add_address, transactions.payment , transaction_status.name, order_types.name, profile.phone_number LIMIT 100

	`

	row, err := db.Query(context.Background(), sql, id)

	fmt.Println(err)

	if err != nil {
		return models.TransactionDetailJoin{}, err
	}

	transaction, err := pgx.CollectOneRow(row, pgx.RowToStructByPos[models.TransactionDetailJoin])

	if err != nil {
		return models.TransactionDetailJoin{}, err
	}

	return transaction, nil
}

func FindTransactionProductById(id int) ([]models.TransactionProduct, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `
	SELECT no_order, products.title, transaction_details.quantity, product_variants.name, product_sizes.name, order_types.name, products.price
	FROM transactions
	INNER JOIN transaction_details on transactions.transaction_detail_id = transaction_details.id
	INNER JOIN products on transaction_details.product_id = products.id
	INNER JOIN product_sizes on transaction_details.product_size_id = product_sizes.id
	INNER JOIN product_variants on transaction_details.variant_id = product_variants.id
	INNER JOIN order_types on transactions.order_type_id = order_types.id
    WHERE no_order = $1
	`

	row, err := db.Query(context.Background(), sql, id)

	fmt.Println(err)

	if err != nil {
		return []models.TransactionProduct{}, err
	}

	transaction, err := pgx.CollectRows(row, pgx.RowToStructByPos[models.TransactionProduct])

	if err != nil {
		return []models.TransactionProduct{}, err
	}

	return transaction, nil
}

func FindTransactionByUserId(id int) ([]models.TransactionJoin, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `
		SELECT transactions.no_order, transaction_status.name as order_type, array_agg(transaction_details.quantity) as quantity,
		array_agg(products.price) as price, product_images.image FROM transactions
		INNER JOIN transaction_details ON transactions.transaction_detail_id = transaction_details.id
		INNER JOIN products ON transaction_details.id = products.id
		INNER JOIN transaction_status ON transactions.transaction_status_id = transaction_status.id
		INNER JOIN product_images ON  product_images.product_id  = products.id
		WHERE transactions.user_id = $1
    	GROUP BY transactions.no_order, transaction_status.name, product_images.image
	`

	row, err := db.Query(context.Background(), sql, id)

	if err != nil {
		log.Println(err)
		return []models.TransactionJoin{}, err
	}

	transaction, err := pgx.CollectRows(row, pgx.RowToStructByName[models.TransactionJoin])

	if err != nil {
		return []models.TransactionJoin{}, err
	}

	return transaction, nil
}

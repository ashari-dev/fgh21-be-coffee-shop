package repository

import (
	"RGT/konis/lib"
	"RGT/konis/models"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func FindTransactionDetailById(id int) (models.TransactionDetailJoin, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `SELECT  transactions.no_order, transactions.add_full_name, transactions.add_address, transactions.payment , transaction_status.name AS transaction_status, transaction_details.quantity, order_types.name AS order_type, profile.phone_number, products.title, product_variants.name as variant, product_sizes.name as size
	FROM transaction_details
	INNER JOIN transactions ON transaction_details.transaction_id=transactions.id
	INNER JOIN transaction_status on transactions.transaction_status_id = transaction_status.id
	INNER JOIN order_types on transactions.order_type_id = order_types.id
	INNER JOIN profile on transactions.user_id = profile.user_id
	INNER JOIN products on transaction_details.product_id = products.id
	INNER JOIN product_variants on transaction_details.variant_id = transaction_details.variant_id
	INNER JOIN product_sizes on transaction_details.product_size_id = product_sizes.id		
	`

	row, err := db.Query(context.Background(), sql, id)

	fmt.Println(err)

	if err != nil {
		return models.TransactionDetailJoin{}, err
	}

	user, err := pgx.CollectOneRow(row, pgx.RowToStructByPos[models.TransactionDetailJoin])

	if err != nil {
		return models.TransactionDetailJoin{}, err
	}

	return user, nil
}

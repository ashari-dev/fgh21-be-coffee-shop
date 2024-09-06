package repository

import (
	"RGT/konis/lib"
	"RGT/konis/models"
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

func CreateTransaction(data models.Transaction) (models.Transaction, error) {
	db := lib.DB()
	defer db.Close(context.Background())
	fmt.Println(data)
	sql := `
		INSERT INTO transactions (no_order, add_full_name, add_email,
		add_address, payment, user_id, transaction_detail_id,
		order_type_id, transaction_status_id) VALUES 
		($1, $2, $3, $4, $5, $6,$7, $8, $9) RETURNING *
	`

	row, err := db.Query(context.Background(), sql, data.NoOrder, data.AddFullName, data.AddEmail, data.AddAddress, data.Payment, data.UserId, data.TransactionDetail, data.OrderTypeId, data.TransactionStatusId)

	if err != nil {
		return models.Transaction{}, err
	}

	transaction, err := pgx.CollectOneRow(row, pgx.RowToStructByPos[models.Transaction])

	if err != nil {
		return models.Transaction{}, err
	}

	return transaction, err
}


func FindAllTransactions(search string, page int, limit int) ([]models.AllTransactionForAdmin, int) {
	db := lib.DB()
	defer db.Close(context.Background())
	offset := 0
	if page > 1 {
		offset = (page - 1) * limit
	}

	sql := `SELECT transactions.id, transactions.no_order, transaction_details.quantity, (products.price * transaction_details.quantity) AS total_price, products.title, transaction_status.name as order_status  FROM transactions
		INNER JOIN transaction_details ON transaction_detail_id = transaction_details.id
		INNER JOIN products ON transaction_details.product_id = products.id
		INNER JOIN transaction_status ON transaction_status_id = transaction_status.id
		WHERE products.title ilike '%' || $1 || '%'
		limit $2 offset $3`

	rows, _ := db.Query(context.Background(),
		sql, search, limit, offset,
	)
	count := TotalTransactions(search)

	transaction, err := pgx.CollectRows(rows, pgx.RowToStructByPos[models.AllTransactionForAdmin])
	if err != nil {
		return []models.AllTransactionForAdmin{}, count
	}

	return transaction, count
}

func TotalTransactions(search string) int {
	db := lib.DB()
	defer db.Close(context.Background())
	inputSQL := `select count(no_order) as "total" FROM transactions
		INNER JOIN transaction_details ON transaction_detail_id = transaction_details.id
		INNER JOIN products ON transaction_details.product_id = products.id
		INNER JOIN transaction_status ON transaction_status_id = transaction_status.id
		WHERE products.title ilike '%' || $1 || '%'`
	rows := db.QueryRow(context.Background(), inputSQL, search)
	var result int
	rows.Scan(
		&result,
	)
	return result
}

func FindTransactionsByStatusId(search int, page int, limit int) ([]models.AllTransactionForAdmin, int) {
	db := lib.DB()
	defer db.Close(context.Background())
	offset := 0
	if page > 1 {
		offset = (page - 1) * limit
	}

	sql := `SELECT transactions.id, transactions.no_order, transaction_details.quantity, (products.price * transaction_details.quantity) AS total_price, products.title, transaction_status.name as order_status  FROM transactions
		INNER JOIN transaction_details ON transaction_detail_id = transaction_details.id
		INNER JOIN products ON transaction_details.product_id = products.id
		INNER JOIN transaction_status ON transaction_status_id = transaction_status.id
		WHERE transaction_status_id = $1
		limit $2 offset $3`

	rows, _ := db.Query(context.Background(),
		sql, search, limit, offset,
	)
	count := TotalTransactionsByStatusId(search)

	transaction, err := pgx.CollectRows(rows, pgx.RowToStructByPos[models.AllTransactionForAdmin])
	if err != nil {
		return []models.AllTransactionForAdmin{}, count
	}

	return transaction, count
}

func TotalTransactionsByStatusId(search int) int {
	db := lib.DB()
	defer db.Close(context.Background())
	inputSQL := `select count(no_order) as "total" FROM transactions
		INNER JOIN transaction_details ON transaction_detail_id = transaction_details.id
		INNER JOIN products ON transaction_details.product_id = products.id
		INNER JOIN transaction_status ON transaction_status_id = transaction_status.id
		WHERE transaction_status_id = $1`
	rows := db.QueryRow(context.Background(), inputSQL, search)
	var result int
	rows.Scan(
		&result,
	)
	return result
}

func FindOneTransactionById(id int) (models.AllTransactionForAdmin, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	inputSQL := `SELECT transactions.id, transactions.no_order, transaction_details.quantity, (products.price * transaction_details.quantity) AS total_price, products.title, transaction_status.name as order_status  FROM transactions
		INNER JOIN transaction_details ON transaction_detail_id = transaction_details.id
		INNER JOIN products ON transaction_details.product_id = products.id
		INNER JOIN transaction_status ON transaction_status_id = transaction_status.id
		WHERE transactions.id = $1`

	row, _ := db.Query(context.Background(), inputSQL, id)

	result, err := pgx.CollectOneRow(row, pgx.RowToStructByPos[models.AllTransactionForAdmin])

	if err != nil {
		return models.AllTransactionForAdmin{}, err
	}

	return result, err
}

func DeleteTransaction(id int) (models.AllTransactionForAdmin, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	result, _ := FindOneTransactionById(id)

	sql := `DELETE FROM transactions WHERE transactions.id=$1;`

	query, err := db.Exec(context.Background(), sql, id)

	if err != nil {
		return models.AllTransactionForAdmin{}, err
	}

	if query.RowsAffected() == 0 {
		return models.AllTransactionForAdmin{}, err
  }
 return result , err
}

func EditTransactionStatus(data models.Transaction, id int) (models.Transaction, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `UPDATE transactions SET "transaction_status_id"=$1 WHERE no_order=$2 returning id,transaction_status_id`

	query := db.QueryRow(context.Background(), sql, data.TransactionStatusId ,id)
	
	var result models.Transaction
	err := query.Scan(
		&result.Id,
		&result.TransactionStatusId,
	)
	
	if err != nil {
		log.Println(err)
		return models.Transaction{}, err
	}

	return result, err
}

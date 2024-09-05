package repository

import (
	"RGT/konis/lib"
	"RGT/konis/models"
	"context"

	"github.com/jackc/pgx/v5"
)

func FindTransactionStatusById(id int) (models.TransactionStatus, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `SELECT * from transaction_status where "id" = $1`

	query, err := db.Query(context.Background(), sql, id)

	if err != nil {
		return models.TransactionStatus{}, err
	}

	status, err := pgx.CollectOneRow(query, pgx.RowToStructByPos[models.TransactionStatus])

	if err != nil {
		return models.TransactionStatus{}, err
	}

	return status, err
}

func FindAllTransactionStatus() ([]models.TransactionStatus, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	rows, err := db.Query(context.Background(),
		`select * from "transaction_status"`,
	)

	if err != nil {
		return []models.TransactionStatus{}, err
	}

	status, err := pgx.CollectRows(rows, pgx.RowToStructByPos[models.TransactionStatus])
	if err != nil {
		return []models.TransactionStatus{}, err
	}
	return status, err
}


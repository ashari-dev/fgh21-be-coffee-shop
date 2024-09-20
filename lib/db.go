package lib

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func DB() *pgx.Conn {

	conn, err := pgx.Connect(context.Background(), "postgresql://postgres:123@3.25.204.209:5432/konis_caffee?sslmode=disable")

	if err != nil {
		fmt.Println(err)
	}
	return conn
}

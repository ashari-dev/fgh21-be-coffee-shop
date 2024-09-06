package models

type Transaction struct {
	Id                  int
	NoOrder             int
	AddFullName         string
	AddEmail            string
	AddAddress          string
	Payment             string
	UserId              int
	TransactionDetail   int
	OrderTypeId         int
	TransactionStatusId int
}

type TransactionJoin struct {
	NoOrder    int    `db:"no_order"`
	OrderTypes string `db:"order_type"`
	Image      string `db:"image"`
	Quantity   []int  `db:"quantity"`
	Price      []int  `db:"price"`
}

type AllTransactionForAdmin struct {
	Id          int    `json:"id"`
	NoOrder     int    `json:"noOrder" db:"no_order"`
	Quantity    int    `json:"quantity" db:"quantity"`
	Price       int    `json:"price" db:"price"`
	Title       string `json:"title" db:"title"`
	OrderStatus string `json:"orderStatus" db:"order_status"`
}
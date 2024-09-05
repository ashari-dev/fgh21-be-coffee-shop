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
	NoOrder    int
	OrderTypes string
	Image string
	Quantity   []int
	Price      []int
}

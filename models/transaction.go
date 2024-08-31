package models

type Transaction struct {
	Id int
    NoOrder int
    AddFullName string
    AddEmail string
    AddAddress string
    Payment string
    UserId int
    OrderTypeId int
    TransactionStatusId int
}
package models

type TransactionDetailJoin struct {
	NoOrder           int    `json:"noOrder"`
	FullName          string `json:"fullName"`
	Address           string `json:"address"`
	Payment           string `json:"payment"`
	TransactionStatus string `json:"transactionStatus"`
	Quantity          int    `json:"quantity"`
	OrderType         string `json:"orderType"`
	PhoneNumber       string `json:"phoneNumber"`
	Title             string `json:"title"`
	Variant           string `json:"variant"`
	Size              string `json:"size"`
}

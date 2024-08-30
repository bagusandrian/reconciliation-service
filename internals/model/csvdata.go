package model

import "time"

const (
	Debit TypeTransaction = 1
	Credit
)

type (
	TypeTransaction int
	DataSystemCSV   struct {
		TrxID                 string
		Amount                float64
		Type                  TypeTransaction
		TransactionTimeString string
		TransactionTime       time.Time
	}
	DataBankCSV struct {
		UniqueIdentifier string
		Amount           float64
		DateString       string
		Date             time.Time
	}
)

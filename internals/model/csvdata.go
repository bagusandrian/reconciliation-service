package model

import "time"

const (
	Debit TypeTransaction = 1
	Credit
)

type (
	TypeTransaction int
	DataSystem      struct {
		DataSystemCSV []DataSystemCSV
		TotalData     int64
	}
	DataSystemCSV struct {
		TrxID                 string
		Amount                float64
		Type                  TypeTransaction
		TransactionTimeString string
		TransactionTime       time.Time
		MatchTransaction      bool
		BankID                string
	}
	DataBank struct {
		DataBankCSV map[string][]DataBankCSV
		TotalData   int64
	}
	DataBankCSV struct {
		UniqueIdentifier string
		Amount           float64
		Type             TypeTransaction
		DateString       string
		Date             time.Time
		MatchTransaction bool
		TrxID            string
	}
)

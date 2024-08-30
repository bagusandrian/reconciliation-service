package model

import "time"

type (
	HeaderResponse struct {
		Status         int    `json:"status"`
		Error          string `json:"error,omitempty"`
		ProcessingTime string `json:"processing_time"`
	}
	ReconciliationRequest struct {
		SystemTransactionCSVFilePath  string             `json:"system_transaction_csv_file_path"`
		BankStatements                []BankCSCVFilePath `json:"bank_statements"`
		ReconciliationStartDateString string             `json:"reconciliaton_start_date"`
		ReconciliationStartDate       time.Time
		ReconciliationEndDateString   string `json:"reconciliaton_end_date"`
		ReconciliationEndDate         time.Time
	}
	BankCSCVFilePath struct {
		BankName    string `json:"bank_name"`
		CSVFilePath string `json:"csv_file_path"`
	}
	ReconciliationResponse struct {
		TotalTranscationsProcessed       int64                                 `json:"total_transactions_processed"`
		TotalNumberMatchedTransactions   int64                                 `json:"total_number_matched_transactions"`
		DetailOfMatchedTransactions      map[string]DetailMatchedTransaction   `json:"detail_of_matched_transactions"`
		TotalNumberUnmatchedTransactions int64                                 `json:"total_number_unmatched_transactions"`
		DetailOfUnmatchedTransactions    map[string]DetailUnmatchedTransaction `json:"detail_of_unmatched_transactions"`
		TotalDiscrepanciesAmount         float64                               `json:"total_discrepancies_amount"`
	}
	DetailMatchedTransaction struct {
		TotalNumberMatchedTransactions int64 `json:"total_number_matched_transactions"`
	}
	DetailUnmatchedTransaction struct {
		Info              string `json:"info"`
		DetailTransaction struct {
			TrxID            string  `json:"trx_id,omitempty"`
			UniqueIdentifier string  `json:"unique_identifier,omitempty"`
			Amount           float64 `json:"amount,omitempty"`
			Date             string  `json:"date,omitempty"`
			TransactionTime  string  `json:"transaction_time,omitempty"`
		} `json:"detail_transaction"`
	}
)

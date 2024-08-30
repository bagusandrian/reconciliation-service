package constant

var (
	ValidateHeaderSystemCSV = []string{"trxID", "amount", "type", "transactionTime"}
	ValidateHeaderBankCSV   = []string{"unique_identifier", "amount", "date"}
)

const (
	LenRowSystem = 4
	LenRowBank   = 3
)

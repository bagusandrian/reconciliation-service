package readfile

import "github.com/bagusandrian/reconciliation-service/internals/model"

//go:generate mockery --name=ReadFile --filename=mock_read_file.go --inpackage
type ReadFile interface {
	GetSystemReconciliationCSV(req model.ReconciliationRequest) (resp model.DataSystem, err error)
	GetBankReconciliationCSV(req model.ReconciliationRequest) (resp model.DataBank, err error)
}

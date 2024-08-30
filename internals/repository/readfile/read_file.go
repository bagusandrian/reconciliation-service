package readfile

import "github.com/bagusandrian/reconciliation-service/internals/model"

//go:generate mockery --name=Handler --filename=mock_handler.go --inpackage
type ReadFile interface {
	GetSystemReconciliationCSV(req model.ReconciliationRequest) (resp []model.DataSystemCSV, err error)
	GetBankReconciliationCSV(req model.ReconciliationRequest) (resp map[string][]model.DataBankCSV, err error)
}

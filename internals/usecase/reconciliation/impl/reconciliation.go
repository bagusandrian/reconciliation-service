package impl

import (
	"context"
	"log"

	"github.com/bagusandrian/reconciliation-service/internals/model"
)

func (u *usecase) ReconciliationComparition(ctx context.Context, req model.ReconciliationRequest) (resp model.ReconciliationResponse, err error) {
	// get csv file system
	dataSystemCSV, err := u.readFile.GetSystemReconciliationCSV(req)
	log.Println(dataSystemCSV)
	dataBankCSV, err := u.readFile.GetBankReconciliationCSV(req)
	log.Println(dataBankCSV)
	return resp, err
}

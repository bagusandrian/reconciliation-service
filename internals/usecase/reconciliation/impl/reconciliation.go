package impl

import (
	"context"
	"log"

	"github.com/bagusandrian/reconciliation-service/internals/model"
)

func (u *usecase) ReconciliationComparition(ctx context.Context, req model.ReconciliationRequest) (resp model.ReconciliationResponse, err error) {
	// get csv file system
	a, err := u.readFile.GetSystemReconciliationCSV(req)
	if err != nil {
		return resp, err
	}
	b, err := u.readFile.GetBankReconciliationCSV(req)
	if err != nil {
		return resp, err
	}
	log.Println(len(a))
	for bankName, v := range b {
		log.Println(bankName, ": ", len(v))
	}
	return resp, err
}

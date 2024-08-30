package reconciliation

import (
	"context"

	"github.com/bagusandrian/reconciliation-service/internals/model"
)

//go:generate mockery --name=Usecase --filename=mock_usecase.go --inpackage
type Usecase interface {
	// reconciliation usecase
	ReconciliationComparition(ctx context.Context, req model.ReconciliationRequest) (resp model.ReconciliationResponse, err error)
}

package dummy

import (
	"context"

	"github.com/bagusandrian/reconciliation-service/internals/model"
)

//go:generate mockery --name=Usecase --filename=mock_usecase.go --inpackage
type Usecase interface {
	GetDummy(ctx context.Context, request model.GetDummyRequest) (response model.GetDummyResponse, err error)
}

package impl

import (
	httpframework "github.com/bagusandrian/reconciliation-service/internals/handler/http"
	"github.com/bagusandrian/reconciliation-service/internals/model"
	ucReconciliation "github.com/bagusandrian/reconciliation-service/internals/usecase/reconciliation"
)

type handler struct {
	config  *model.Config
	usecase ucReconciliation.Usecase
}

func New(
	cfg *model.Config, ucReconciliation ucReconciliation.Usecase,
) httpframework.Handler {

	// init repository
	h := &handler{
		config:  cfg,
		usecase: ucReconciliation,
	}
	return h
}

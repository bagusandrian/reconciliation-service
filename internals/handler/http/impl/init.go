package impl

import (
	httpframework "github.com/bagusandrian/reconciliation-service/internals/handler/http"
	"github.com/bagusandrian/reconciliation-service/internals/model"
	ucDummy "github.com/bagusandrian/reconciliation-service/internals/usecase/dummy"
)

type handler struct {
	config  *model.Config
	usecase ucDummy.Usecase
}

func New(
	cfg *model.Config, ucDummy ucDummy.Usecase,
) httpframework.Handler {

	// init repository
	h := &handler{
		config:  cfg,
		usecase: ucDummy,
	}
	return h
}

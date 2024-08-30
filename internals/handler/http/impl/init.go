package impl

import (
	"github.com/bagusandrian/reconciliation-service/internals/config"
	httpframework "github.com/bagusandrian/reconciliation-service/internals/handler/http"
	ucDummy "github.com/bagusandrian/reconciliation-service/internals/usecase/dummy"
)

type handler struct {
	config  *config.Config
	usecase ucDummy.Usecase
}

func New(
	cfg *config.Config, ucDummy ucDummy.Usecase,
) httpframework.Handler {

	// init repository
	h := &handler{
		config:  cfg,
		usecase: ucDummy,
	}
	return h
}

package impl

import (
	"github.com/bagusandrian/reconciliation-service/internals/model"
	"github.com/bagusandrian/reconciliation-service/internals/repository/readfile"
	"github.com/bagusandrian/reconciliation-service/internals/usecase/reconciliation"
)

type usecase struct {
	cfg      *model.Config
	readFile readfile.ReadFile
}

func New(cfg *model.Config, readFile readfile.ReadFile) reconciliation.Usecase {
	return &usecase{
		cfg:      cfg,
		readFile: readFile,
	}

}

package impl

import (
	"github.com/bagusandrian/reconciliation-service/internals/model"
	repository "github.com/bagusandrian/reconciliation-service/internals/repository/db"
	"github.com/bagusandrian/reconciliation-service/internals/usecase/dummy"
)

type usecase struct {
	cfg  *model.Config
	repo repository.DB
}

func New(cfg *model.Config, repo repository.DB) dummy.Usecase {
	return &usecase{
		cfg:  cfg,
		repo: repo,
	}

}

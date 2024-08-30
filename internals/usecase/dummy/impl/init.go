package impl

import (
	"github.com/bagusandrian/reconciliation-service/internals/config"
	repository "github.com/bagusandrian/reconciliation-service/internals/repository/db"
	"github.com/bagusandrian/reconciliation-service/internals/usecase/dummy"
)

type usecase struct {
	cfg  *config.Config
	repo repository.DB
}

func New(cfg *config.Config, repo repository.DB) dummy.Usecase {
	return &usecase{
		cfg:  cfg,
		repo: repo,
	}

}

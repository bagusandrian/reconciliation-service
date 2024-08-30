package impl

import (
	"github.com/bagusandrian/reconciliation-service/internals/model"
	"github.com/bagusandrian/reconciliation-service/internals/repository/readfile"
)

type repoReadFile struct {
	cfg *model.Config
}

func New(cfg *model.Config) readfile.ReadFile {
	return &repoReadFile{
		cfg: cfg,
	}
}

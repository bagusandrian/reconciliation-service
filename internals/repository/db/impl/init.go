package impl

import (
	"github.com/bagusandrian/reconciliation-service/internals/model"
	"github.com/bagusandrian/reconciliation-service/internals/repository/db"

	"database/sql"
)

type repoDB struct {
	cfg *model.Config
	db  *sql.DB
}

func New(cfg *model.Config) db.DB {
	return &repoDB{
		cfg: cfg,
		db:  &sql.DB{},
	}
}

package fundraising_storage

import (
	"github.com/jmoiron/sqlx"
	"mono-tracker/internal/domain/fundraising"
)

type Storage struct {
	db *sqlx.DB
}

func NewStorage(db *sqlx.DB) fundraising.IFundraisingStorage {
	return &Storage{
		db: db,
	}
}

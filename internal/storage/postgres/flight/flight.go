package flight

import "github.com/jmoiron/sqlx"

type Flight struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Flight {
	return &Flight{
		db: db,
	}
}

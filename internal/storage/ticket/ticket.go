package ticket

import "github.com/jmoiron/sqlx"

type Ticket struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Ticket {
	return &Ticket{
		db: db,
	}
}

package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(UpTicket, DownTicket)
}

func UpTicket(ctx context.Context, tx *sql.Tx) error {
	query := `CREATE TABLE IF NOT EXISTS tickets (
    	id SERIAL PRIMARY KEY,
    	cost INTEGER NOT NULL,
    	state_number TEXT REFERENCES transport(state_number) NOT NULL
	);`

	_, err := tx.ExecContext(ctx, query)
	return err
}

func DownTicket(ctx context.Context, tx *sql.Tx) error {
	query := `DROP TABLE IF EXISTS tickets;`
	_, err := tx.ExecContext(ctx, query)
	return err
}

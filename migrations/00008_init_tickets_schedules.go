package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(UpTicketsSchedules, DownTicketsSchedules)
}

func UpTicketsSchedules(ctx context.Context, tx *sql.Tx) error {
	query := `CREATE TABLE IF NOT EXISTS tickets_flights (
    	ticket_id INTEGER REFERENCES tickets(id) ON DELETE CASCADE,
    	flight_id INTEGER REFERENCES flights(id) ON DELETE CASCADE
	);`

	_, err := tx.ExecContext(ctx, query)
	return err
}

func DownTicketsSchedules(ctx context.Context, tx *sql.Tx) error {
	query := `DROP TABLE IF EXISTS tickets_flights`
	_, err := tx.ExecContext(ctx, query)
	return err
}

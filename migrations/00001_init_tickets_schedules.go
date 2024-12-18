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
	query := `CREATE TABLE IF NOT EXISTS tickets_schedules (
    	ticket_id INTEGER REFERENCES tickets(id) ON DELETE CASCADE,
    	schedule_id INTEGER REFERENCES schedules(id) ON DELETE CASCADE
	);`

	_, err := tx.ExecContext(ctx, query)
	return err
}

func DownTicketsSchedules(ctx context.Context, tx *sql.Tx) error {
	query := `DROP TABLE IF EXISTS tickets_schedules`
	_, err := tx.ExecContext(ctx, query)
	return err
}

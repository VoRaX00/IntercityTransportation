package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(UpSchedule, DownSchedule)
}

func UpSchedule(ctx context.Context, tx *sql.Tx) error {
	query := `CREATE TABLE IF NOT EXISTS schedules (
    	id SERIAL PRIMARY KEY,
    	from_id INTEGER REFERENCES place(id) NOT NULL,
    	to_id INTEGER REFERENCES place(id) NOT NULL,
    	departure TIMESTAMP NOT NULL,
    	arrival TIMESTAMP NOT NULL
	);`

	_, err := tx.ExecContext(ctx, query)
	return err
}

func DownSchedule(ctx context.Context, tx *sql.Tx) error {
	query := `DROP TABLE IF EXISTS schedules;`
	_, err := tx.ExecContext(ctx, query)
	return err
}

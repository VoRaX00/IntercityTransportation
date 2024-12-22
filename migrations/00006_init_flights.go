package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(UpFlight, DownSchedule)
}

func UpFlight(ctx context.Context, tx *sql.Tx) error {
	query := `CREATE TABLE IF NOT EXISTS flights (
    	id SERIAL PRIMARY KEY,
    	from_id INTEGER REFERENCES places(id) NOT NULL,
    	to_id INTEGER REFERENCES places(id) NOT NULL,
    	departure TIMESTAMP NOT NULL,
    	arrival TIMESTAMP NOT NULL,
    	state_number TEXT REFERENCES buses(state_number) NOT NULL
	);`

	_, err := tx.ExecContext(ctx, query)
	return err
}

func DownSchedule(ctx context.Context, tx *sql.Tx) error {
	query := `DROP TABLE IF EXISTS flights;`
	_, err := tx.ExecContext(ctx, query)
	return err
}

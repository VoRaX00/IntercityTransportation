package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(UpTransport, DownTransport)
}

func UpTransport(ctx context.Context, tx *sql.Tx) error {
	query := `CREATE TABLE IF NOT EXISTS buses (
    	state_number VARCHAR(10) PRIMARY KEY,
    	model TEXT REFERENCES models(model) NOT NULL
	);`

	_, err := tx.ExecContext(ctx, query)
	return err
}

func DownTransport(ctx context.Context, tx *sql.Tx) error {
	query := `DROP TABLE IF EXISTS buses;`
	_, err := tx.ExecContext(ctx, query)
	return err
}

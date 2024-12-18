package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(UpModels, DownModels)
}

func UpModels(ctx context.Context, tx *sql.Tx) error {
	query := `CREATE TABLE IF NOT EXISTS models (
    	model TEXT PRIMARY KEY,
    	count_places INTEGER NOT NULL
	);`

	_, err := tx.ExecContext(ctx, query)
	return err
}

func DownModels(ctx context.Context, tx *sql.Tx) error {
	query := `DROP TABLE IF EXISTS models;`
	_, err := tx.ExecContext(ctx, query)
	return err
}

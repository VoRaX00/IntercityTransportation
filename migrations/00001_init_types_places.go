package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(UpTypesPlaces, DownTypesPlaces)
}

func UpTypesPlaces(ctx context.Context, tx *sql.Tx) error {
	query := `CREATE TABLE IF NOT EXISTS types_places (
    	id SERIAL PRIMARY KEY,
    	type_place TEXT NOT NULL
	);`

	_, err := tx.ExecContext(ctx, query)
	return err
}

func DownTypesPlaces(ctx context.Context, tx *sql.Tx) error {
	query := `DROP TABLE IF EXISTS types_places;`
	_, err := tx.ExecContext(ctx, query)
	return err
}

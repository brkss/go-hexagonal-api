package postgres

import (
	"context"
	"embed"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/brkss/dextrace-server/internal/adapter/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

// migrationsFS is a filesystem that embeds the migrations folder

//go:embed migrations/*.sql
var migrationsFS embed.FS

/**
* DB is a wrapper for PostgreSQl database connection 
* that uses pgxpool as database driver.
* It also holds a refrence to squirrel.StatemenetBuilderType
* Which is used to build SQL queries that compatible with postgreSQL
*/

type DB struct {
	*pgxpool.Pool
	QueryBuilder *squirrel.StatementBuilderType
	url string
}

// New create a nw PostgreSQL database instance 
func New(ctx context.Context, config *config.DB)(*DB, error) {
	url := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable",
		config.Connection,
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
	)

	db, err := pgxpool.New(ctx, url)
	if err != nil {
		return nil, err
	}
	
	err = db.Ping(ctx)
	if err != nil {
		return nil, err;
	}

	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	return &DB {
		db,
		&psql,
		url, 
	}, nil
}


// Migrate runs the database migration 
func (db *DB) Migrate() error {
	driver, err := iofs.New(migrationsFS, "migrations")
	if err != nil {
		return err;
	}

	migrations, err := migrate.NewWithSourceInstance("iofs", driver, db.url)
	if err != nil {
		return err;
	}

	err = migrations.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err;
	}

	return nil;
}

// ErrorCode returns the error code of the given error
func (db *DB) ErrorCode(err error) string {
	pgErr := err.(*pgconn.PgError)
	return pgErr.Code
}

// Close closes the database connection
func (db *DB) Close() {
	db.Pool.Close()
}
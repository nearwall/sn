package postgres

import (
	"context"
	"errors"
	"net/url"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/uptrace/opentelemetry-go-extra/otelsql"
	"github.com/uptrace/opentelemetry-go-extra/otelsqlx"
	semconv "go.opentelemetry.io/otel/semconv/v1.30.0"
)

func connect(config Config) (*sqlx.DB, error) {
	sslMode := "require"
	if config.DisableTLS {
		sslMode = "disable"
	}

	query := make(url.Values)

	query.Set("sslmode", sslMode)
	query.Set("timezone", timezone)

	if config.Schema != "" {
		query.Set("search_path", config.Schema)
	}

	url := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(config.User, config.Password),
		Host:     config.Host,
		Path:     config.DBName,
		RawQuery: query.Encode(),
	}

	db, err := otelsqlx.Open("pgx", url.String(), otelsql.WithAttributes(semconv.DBSystemNamePostgreSQL))
	if err != nil {
		return nil, err
	}

	if config.MaxIdleConns == 0 {
		db.SetMaxIdleConns(DefaultMaxIdleConns)
	} else {
		db.SetMaxIdleConns(config.MaxIdleConns)
	}

	if config.MaxOpenConns == 0 {
		db.SetMaxOpenConns(DefaultMaxOpenConns)
	} else {
		db.SetMaxOpenConns(config.MaxOpenConns)
	}

	if err = ping(context.Background(), db); err != nil {
		return nil, errors.New(
			"failed to ping the database: " +
				config.Host +
				". Please check DB existence, username and his password, Postgres status")
	}

	return db, nil
}

func ping(ctx context.Context, db *sqlx.DB) error {
	if _, ok := ctx.Deadline(); ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, pingTimeout)
		defer cancel()
	}

	for attempts := 1; ; attempts++ {
		err := db.Ping()
		if err == nil {
			break
		}

		if attempts == maxPingAttempts {
			return err
		}

		time.Sleep(time.Duration(attempts) * sleepTimeBetweenAttempts * time.Millisecond)
	}

	// Simple query request for checking connection
	var result bool
	return db.QueryRowContext(ctx, "SELECT true").Scan(&result)
}

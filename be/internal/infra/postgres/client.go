package postgres

import (
	"context"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"

	// The blank import is used to ensure that the "github.com/jackc/pgx/v5/stdlib" package's init functions are executed, which registers the PostgreSQL driver.
	_ "github.com/jackc/pgx/v5/stdlib"
)

const (
	timezone                 = "utc"
	pingTimeout              = 1000 * time.Millisecond
	maxPingAttempts          = 10
	sleepTimeBetweenAttempts = 100
)

const (
	DefaultMaxOpenConns = 5
	DefaultMaxIdleConns = 2
)

type (
	Config struct {
		// DB username
		User string
		// DB password
		Password string
		// DB host including port (e.g. "localhost:5432")
		Host string
		// DB name
		DBName string
		// Data schema (using as search_path)
		Schema       string
		MaxIdleConns int
		MaxOpenConns int
		DisableTLS   bool

		// DB name for DB creation if not exists (default: postgres)
		DefaultDBName string
	}

	Client struct {
		*sqlx.DB
	}
)

// NewClient creates DB if it does not exist and returns Postgres client
func NewClient(config Config) (*Client, error) {
	if err := createDBIfNotExists(config); err != nil {
		return nil, errors.New("failed to connect to the default database: " + err.Error())
	}

	db, err := connect(config)
	if err != nil {
		return nil, err
	}

	// logger.Log().Infof(context.Background(), "Postgres client successfully connect to %v", config.Host)

	return &Client{db}, nil
}

func (client *Client) Ping(ctx context.Context) error {
	return ping(ctx, client.DB)
}

func (client *Client) Close() error {
	return client.DB.Close()
}

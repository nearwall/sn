package inject

import (
	"context"
	"encoding/hex"

	"sn/internal/infra/postgres"
	repository "sn/internal/repository/user"
	"sn/internal/service"

	"github.com/google/wire"
	"github.com/urfave/cli/v3"
)

// wire Set for loading the services.
var serviceSet = wire.NewSet( // nolint
	providePostgresClient,
	service.NewUserService,
	provideJWTServiceConfig,
	service.NewJWTService,
	repository.NewStore,
)

func providePostgresClient(ctx context.Context, cmd *cli.Command) (*postgres.Client, error) {
	return postgres.NewClient(postgres.Config{
		User:       cmd.String("postgres-user"),
		Password:   cmd.String("postgres-password"),
		Host:       cmd.String("postgres-host"),
		DBName:     cmd.String("postgres-db-name"),
		DisableTLS: cmd.Bool("postgres-disable-tls"),

		MaxOpenConns: 3,
		MaxIdleConns: 1,
	})
}

func provideJWTServiceConfig(ctx context.Context, cmd *cli.Command) (service.TokenServiceConfig, error) {
	key, err := hex.DecodeString(cmd.String("jwt-key"))

	return service.TokenServiceConfig{
		Key: key,
	}, err
}

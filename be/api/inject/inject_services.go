package inject

import (
	"context"
	"encoding/hex"
	"time"

	"sn/internal/core"
	"sn/internal/infra/postgres"
	acc_stg "sn/internal/repository/account"
	info_stg "sn/internal/repository/info"
	"sn/internal/service"

	"github.com/google/wire"
	"github.com/urfave/cli/v3"
)

// wire Set for loading the services.
var serviceSet = wire.NewSet( // nolint
	providePostgresClient,
	service.NewAccountService,
	provideJWTServiceConfig,
	service.NewTokenService,
	info_stg.NewInfoStore,
	acc_stg.NewAccountStore,
	service.NewAuthService,
	providePasswordServiceConfig,
	service.NewPasswordService,
	provideSessionServiceConfig,
	service.NewSessionService,
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

func providePasswordServiceConfig(_ctx context.Context, _cmd *cli.Command) (service.PasswordServiceConfig, error) {
	return service.PasswordServiceConfig{
		HashPepper: "",
		// FixMe: add cli flag
		HashAlgorithm: core.HashAlgoDebugBytesSum,
	}, nil
}

func provideSessionServiceConfig(_ctx context.Context, _cmd *cli.Command) (service.SessionServiceConfig, error) {
	// FixMe: add cli flag
	return service.SessionServiceConfig{
		SessionDuration: 32 * time.Hour,
	}, nil
}

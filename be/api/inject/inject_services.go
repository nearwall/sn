package inject

import (
	"context"
	"encoding/hex"
	"fmt"

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
	key, err := hex.DecodeString(cmd.String("token-secret-key"))

	return service.TokenServiceConfig{
		Key:                 key,
		AccessTokenLifespan: cmd.Duration("access-token-lifespan"),
	}, err
}

func providePasswordServiceConfig(_ctx context.Context, cmd *cli.Command) (service.PasswordServiceConfig, error) {
	var pwdAlgoID core.PwdHashAlgorithm
	switch ID := cmd.Uint8("password-hash-algorithm-id"); ID {
	case uint8(core.HashAlgoArgon2ID):
		pwdAlgoID = core.HashAlgoArgon2ID
	default:
		return service.PasswordServiceConfig{}, fmt.Errorf("unknown hash algorithm ID: %d", ID)
	}

	var hashPepper *string
	if pepper := cmd.String("password-hash-pepper"); len(pepper) == 0 {
		hashPepper = &pepper
	}

	return service.PasswordServiceConfig{
		HashPepper:    hashPepper,
		HashAlgorithm: pwdAlgoID,
	}, nil
}

func provideSessionServiceConfig(_ctx context.Context, cmd *cli.Command) (service.SessionServiceConfig, error) {
	return service.SessionServiceConfig{
		// ToDo: add special cli flag after adding refresh token
		SessionDuration: cmd.Duration("access-token-lifespan"),
	}, nil
}

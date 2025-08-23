package inject

import (
	"context"
	"encoding/hex"
	"fmt"

	"sn/internal/core"
	"sn/internal/infra/postgres"
	acc_stg "sn/internal/repository/account"
	info_stg "sn/internal/repository/info"
	"sn/internal/service/account"
	"sn/internal/service/auth"
	"sn/internal/service/password"
	"sn/internal/service/password/argon2id"
	"sn/internal/service/password/summ"
	"sn/internal/service/session"
	"sn/internal/service/token"

	"github.com/google/wire"
	"github.com/urfave/cli/v3"
)

// wire Set for loading the services.
var serviceSet = wire.NewSet( // nolint
	providePostgresClient,
	account.NewAccountService,
	provideJWTServiceConfig,
	token.NewTokenService,
	info_stg.NewInfoStore,
	acc_stg.NewAccountStore,
	auth.NewAuthService,
	providePasswordServiceConfig,
	providePasswordService,
	provideSessionServiceConfig,
	session.NewSessionService,
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

func provideJWTServiceConfig(ctx context.Context, cmd *cli.Command) (token.TokenServiceConfig, error) {
	key, err := hex.DecodeString(cmd.String("token-secret-key"))
	if err != nil {
		return token.TokenServiceConfig{}, fmt.Errorf("fail to parse 'token-secret-key' as hex bytes string: %w", err)
	}

	return token.TokenServiceConfig{
		Key:                 key,
		AccessTokenLifespan: cmd.Duration("access-token-lifespan"),
	}, nil
}

func providePasswordServiceConfig(_ctx context.Context, cmd *cli.Command) (password.Config, error) {
	var pwdAlgoID core.PwdHashAlgorithm
	switch ID := cmd.Uint8("password-hash-algorithm-id"); ID {
	case uint8(core.HashAlgoArgon2ID):
		pwdAlgoID = core.HashAlgoArgon2ID
	default:
		return password.Config{}, fmt.Errorf("unknown hash algorithm ID: %d", ID)
	}

	var hashPepper *string
	if pepper := cmd.String("password-hash-pepper"); len(pepper) == 0 {
		hashPepper = &pepper
	}

	return password.Config{
		HashPepper:    hashPepper,
		HashAlgorithm: pwdAlgoID,
	}, nil
}

func providePasswordService(config password.Config) core.PasswordService {
	switch config.HashAlgorithm {
	case core.HashAlgoArgon2ID:
		return argon2id.NewService(
			argon2id.Config{
				Memory:      64 * 1024,
				Iterations:  4,
				Parallelism: 2,
				SaltLength:  16,
				KeyLength:   32,
			},
			nil,
		)
	default:
		return summ.NewService()
	}
}

func provideSessionServiceConfig(_ctx context.Context, cmd *cli.Command) (session.SessionServiceConfig, error) {
	return session.SessionServiceConfig{
		// ToDo: add special cli flag after adding refresh token
		SessionDuration: cmd.Duration("access-token-lifespan"),
	}, nil
}

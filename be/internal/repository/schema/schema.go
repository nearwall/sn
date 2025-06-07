package migrate

import (
	"context"
	"embed"
	"errors"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" //nolint
	_ "github.com/golang-migrate/migrate/v4/source/file"       //nolint
	"github.com/golang-migrate/migrate/v4/source/iofs"

	"sn/internal/infra/logger"
)

var (
	//go:embed migrations/*
	migrationFiles embed.FS
)

func Up(ctx context.Context, databaseURL string) error {
	migrations, err := iofs.New(migrationFiles, "migrations")
	if err != nil {
		logger.Log().Error(ctx, "Fail to read migrations", logger.ErrorLabel, err.Error())
		return err
	}

	migrator, err := migrate.NewWithSourceInstance("iofs", migrations, databaseURL)
	if err != nil {
		logger.Log().Error(ctx, "Fail to create configure migration", logger.ErrorLabel, err)
		return err
	}
	if err = migrator.Up(); !errors.Is(err, migrate.ErrNoChange) {
		logger.Log().Error(ctx, "Fail to migrate to new version(s)", logger.ErrorLabel, err)
		return err
	}

	version, dirty, err := migrator.Version()
	if err != nil {
		logger.Log().Error(ctx, "Fail to obtain version after migration", logger.ErrorLabel, err)
	}

	logger.Log().Info(ctx, "Applied %d version(dirty: %t)", version, dirty)

	return nil
}

func CreateCleanDB(ctx context.Context, databaseURL string, migrationsURL string) error {
	migrator, err := migrate.New(migrationsURL, databaseURL)
	if err != nil {
		return err
	}
	if err = migrator.Down(); !errors.Is(err, migrate.ErrNoChange) {
		return err
	}
	if err = migrator.Up(); !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	version, dirty, err := migrator.Version()
	if err != nil {
		return err
	}

	logger.Log().Info(ctx, "Applied %d version(dirty: %t)", version, dirty)
	return nil
}

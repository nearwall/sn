package sn

import (
	"github.com/urfave/cli/v3"
)

var cmdFlags = []cli.Flag{
	&cli.BoolFlag{
		Name:    "debug",
		Usage:   "Debug",
		Sources: cli.EnvVars("DEBUG"),
		Value:   true,
	},
	&cli.StringFlag{
		Name:    "server-host",
		Usage:   "server host",
		Sources: cli.EnvVars("SERVER_HOST"),
		Value:   "localhost:3000",
	},
	&cli.BoolFlag{
		Name:    "server-debug",
		Usage:   "server debug",
		Sources: cli.EnvVars("SERVER_DEBUG"),
		Value:   true,
	},
	&cli.BoolFlag{
		Name:    "server-profiling",
		Usage:   "server profiling",
		Sources: cli.EnvVars("SERVER_PROFILING"),
		Value:   false,
	},
	&cli.IntFlag{
		Name:    "metrics-port",
		Usage:   "metrics port",
		Sources: cli.EnvVars("METRICS_PORT"),
		Value:   4000,
	},
	&cli.StringFlag{
		Name:    "postgres-user",
		Usage:   "postgres user",
		Sources: cli.EnvVars("POSTGRES_USER"),
		Value:   "postgres",
	},
	&cli.StringFlag{
		Name:    "postgres-password",
		Usage:   "postgres password",
		Sources: cli.EnvVars("POSTGRES_PASSWORD"),
		Value:   "postgres",
	},
	&cli.StringFlag{
		Name:    "postgres-host",
		Usage:   "postgres host",
		Sources: cli.EnvVars("POSTGRES_HOST"),
		Value:   "localhost:5432",
	},
	&cli.StringFlag{
		Name:    "postgres-db-name",
		Usage:   "postgres database name",
		Sources: cli.EnvVars("POSTGRES_DB_NAME"),
		Value:   "sn",
	},
	&cli.BoolFlag{
		Name:    "postgres-disable-tls",
		Usage:   "postgres disable tls",
		Sources: cli.EnvVars("POSTGRES_DISABLE_TLS"),
		Value:   true,
	},
}

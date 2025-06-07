package migrate

import "github.com/urfave/cli/v3"

var cmdFlags = []cli.Flag{
	&cli.StringFlag{
		Name:    "postgres-user",
		Usage:   "postgres user",
		Sources: cli.EnvVars("POSTGRES_USER"),
		Value:   "admin",
	},
	&cli.StringFlag{
		Name:    "postgres-password",
		Usage:   "postgres password",
		Sources: cli.EnvVars("POSTGRES_PASSWORD"),
		Value:   "admin",
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

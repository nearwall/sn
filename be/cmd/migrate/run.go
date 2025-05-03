package migrate

import (
	"context"
	"fmt"
	migrate "sn/internal/repository/schema"

	"github.com/urfave/cli/v3"
)

var Cmd = cli.Command{
	Name:  "migrate",
	Usage: "set actual version migration",
	Flags: cmdFlags,
	OnUsageError: func(_ctx context.Context, _cmd *cli.Command, err error, _isSubcommand bool) error {
		return err
	},
	Action: run,
}

func run(ctx context.Context, cmd *cli.Command) error {
	return migrate.Up(
		ctx,
		fmt.Sprintf(
			"postgres://%s:%s@%s/%s?sslmode=disable",
			cmd.String("postgres-user"),
			cmd.String("postgres-password"),
			cmd.String("postgres-host"),
			cmd.String("postgres-db-name"),
		),
	)
}

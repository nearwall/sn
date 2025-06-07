package main

import (
	"context"
	"os"

	cli "github.com/urfave/cli/v3"

	"sn/cmd/migrate"
	"sn/cmd/sn"
	"sn/internal/infra/logger"
	"sn/version"
)

// @Version 0.0.0
// @Title social_network
// @Description Social network server
func main() {
	var globalFlags = []cli.Flag{
		&cli.BoolFlag{
			Name:    "debug",
			Usage:   "Debug",
			Sources: cli.EnvVars("DEBUG"),
			Value:   true,
		},
		&cli.BoolFlag{
			Name:    "disable-stack-trace",
			Usage:   "Disable stack trace",
			Sources: cli.EnvVars("DISABLE_STACK_TRACE"),
			Value:   true,
		},
	}

	app := &cli.Command{
		Usage: "SN",
		Commands: []*cli.Command{
			&sn.Cmd,
			&migrate.Cmd,
		},
		EnableShellCompletion: true,
		Flags:                 globalFlags,
		Version:               version.Version + " (" + version.GitCommit + ")",
		OnUsageError: func(_ context.Context, _ *cli.Command, err error, _ bool) error {
			return err
		},
		Before: func(ctx context.Context, cmd *cli.Command) (context.Context, error) {
			serviceName := "social_network"

			if cmd.Bool("debug") {
				logger.SetDebugLogger(serviceName, cmd.Bool("disable-stack-trace"))
			} else {
				logger.SetProductionLogger(serviceName)
			}

			return ctx, nil
		},
	}

	appCtx := context.Background()
	if err := app.Run(appCtx, os.Args); err != nil {
		logger.Log().Errorf(appCtx, "application was stopped: %s", err)
	}
}

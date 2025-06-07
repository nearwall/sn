package sn

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/urfave/cli/v3"

	"sn/api/inject"
	"sn/internal/infra/logger"
)

var Cmd = cli.Command{
	Name:  "sn",
	Usage: "social network server",
	Flags: cmdFlags,
	OnUsageError: func(_ context.Context, _cmd *cli.Command, err error, _isSubcommand bool) error {
		return err
	},
	Action: run,
}

// cli.ActionFunc interface
func run(c context.Context, cmd *cli.Command) error {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	signal.Notify(sig, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		select {
		case <-ctx.Done():
			return
		case s := <-sig:
			logger.Log().Infof(ctx, "signal %s received", s.String())
			cancel()
		}
	}()

	app, err := inject.InitializeApplication(cmd, ctx)
	if err != nil {
		logger.Log().Fatalf(ctx, "fail to initialize server: %s", err)
	}

	defer app.PostgresClient.Close()

	go app.RestServer.Run(ctx)

	<-ctx.Done()

	_ = app.RestServer.Shutdown(ctx)

	logger.Log().Debug(ctx, "ctx end received")

	return nil
}

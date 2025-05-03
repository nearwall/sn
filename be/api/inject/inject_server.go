package inject

import (
	"context"

	"sn/api/rest/handlers"

	"github.com/google/wire"
	"github.com/urfave/cli/v3"
)

// wire Set for loading the server.
var serverSet = wire.NewSet( //nolint
	handlers.NewResolver,
	handlers.NewServer,
	provideRestServerConfig,
)

func provideRestServerConfig(ctx context.Context, cmd *cli.Command) (handlers.ServerConfig, error) {
	return handlers.ServerConfig{
		Addr: cmd.String("addr"),
	}, nil
}

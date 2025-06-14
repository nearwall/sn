package inject

import (
	"context"

	"sn/api/rest"
	"sn/api/rest/handlers"
	"sn/api/rest/middleware"

	"github.com/google/wire"
	"github.com/urfave/cli/v3"
)

// wire Set for loading the server.
var serverSet = wire.NewSet( // nolint: unused // it used for wire autogeneration
	handlers.NewResolver,
	rest.NewServer,
	middleware.NewBearerTokenAuth,
	provideRestServerConfig,
)

func provideRestServerConfig(ctx context.Context, cmd *cli.Command) (rest.ServerConfig, error) {
	return rest.ServerConfig{
		Addr: cmd.String("server-host"),
	}, nil
}

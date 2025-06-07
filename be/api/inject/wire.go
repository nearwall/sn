//go:build wireinject
// +build wireinject

package inject

import (
	"context"
	"sn/api"

	"github.com/google/wire"
	"github.com/urfave/cli/v3"
)

func InitializeApplication(c *cli.Command, appCtx context.Context) (api.Container, error) {
	wire.Build(
		serverSet,
		serviceSet,
		api.NewContainer,
	)
	return api.Container{}, nil
}

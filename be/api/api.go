package api

import (
	"sn/api/rest"
	"sn/internal/infra/postgres"
)

type Container struct {
	PostgresClient *postgres.Client
	RestServer     rest.Server
}

// NewContainer creates a new application struct
func NewContainer(
	postgresClient *postgres.Client,
	server rest.Server,
) Container {
	return Container{
		PostgresClient: postgresClient,
		RestServer:     server,
	}
}

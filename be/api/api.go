package api

import (
	"sn/api/rest/handlers"
	"sn/internal/infra/postgres"
)

type Container struct {
	PostgresClient *postgres.Client
	RestServer     handlers.Server
}

// NewContainer: Create a new application struct
func NewContainer(
	postgresClient *postgres.Client,
	server handlers.Server,
) Container {
	return Container{
		PostgresClient: postgresClient,
		RestServer:     server,
	}
}

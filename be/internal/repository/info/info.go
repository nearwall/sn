package info

import (
	"sn/internal/core"
	"sn/internal/infra/postgres"
)

type infoStore struct {
	client *postgres.Client
}

func NewInfoStore(postgresClient *postgres.Client) core.InfoStore {
	return &infoStore{
		client: postgresClient,
	}
}

package info

import (
	"context"
	"time"

	"github.com/google/uuid"

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

// core.InfoStore interface
func (s *infoStore) LinkToAccount(_ context.Context, accountID uuid.UUID, info core.PersonalInfo) error {
	now := time.Now().UTC()
	sql := `INSERT INTO personal_info (
				account_id,
				first_name,
				second_name,
				birth_date,
				biography,
				city,
				updated_at,
				created_at
			) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err := s.client.DB.Exec(sql, accountID, "", now, now, now)

	return err
}

// core.InfoStore interface
func (s *infoStore) ReadInfo(_ context.Context, accountID uuid.UUID) (core.PersonalInfo, error) {
	sql := `SELECT
				first_name,
				second_name,
				birth_date,
				biography,
				city,
				updated_at,
				created_at
			FROM personal_info
			WHERE account_id=$1`

	var raw rawInfo
	if err := s.client.DB.QueryRowx(sql, accountID).StructScan(&raw); err != nil {
		return core.PersonalInfo{}, err
	}

	return core.PersonalInfo{
		FirstName:  raw.firstName,
		SecondName: raw.secondName,
		Birthdate:  raw.birthdate,
		Biography:  raw.biography,
		City:       raw.city,
	}, nil
}

type (
	rawInfo struct {
		firstName  string    `db:"first_name"`
		secondName string    `db:"second_name"`
		birthdate  time.Time `db:"birth_date"`
		biography  string    `db:"biography"`
		city       string    `db:"city"`
	}
)

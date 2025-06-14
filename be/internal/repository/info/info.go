package info

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"sn/internal/core"
	"sn/internal/infra/postgres"
)

type infoStore struct {
	client *postgres.Client
}

type (
	rawInfo struct {
		// fields are capitalized oly for sqlx correct parsing
		FirstName  string    `db:"first_name"`
		SecondName string    `db:"second_name"`
		Birthdate  time.Time `db:"birth_date"`
		Biography  string    `db:"biography"`
		City       string    `db:"city"`
	}
)

func NewInfoStore(postgresClient *postgres.Client) core.InfoStore {
	return &infoStore{
		client: postgresClient,
	}
}

// core.InfoStore interface
func (s *infoStore) LinkToAccount(ctx context.Context, accountID uuid.UUID, info core.PersonalInfo) error {
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

	if _, err := s.client.DB.ExecContext(
		ctx,
		sql,
		accountID,
		info.FirstName,
		info.SecondName,
		info.Birthdate,
		info.Biography,
		info.City,
		now,
		now); err != nil {
		return fmt.Errorf("fail to create an entry in `personal_info` table: %w", err)
	}

	return nil
}

// core.InfoStore interface
func (s *infoStore) ReadInfo(ctx context.Context, accountID uuid.UUID) (core.PersonalInfo, error) {
	sql := `SELECT
				first_name,
				second_name,
				birth_date,
				biography,
				city
			FROM personal_info
			WHERE account_id=$1`

	var raw rawInfo
	if err := s.client.DB.QueryRowxContext(ctx, sql, accountID).StructScan(&raw); err != nil {
		return core.PersonalInfo{}, fmt.Errorf("fail to read an entry from `personal_info` table: %w", err)
	}

	return core.PersonalInfo{
		FirstName:  raw.FirstName,
		SecondName: raw.SecondName,
		Birthdate:  raw.Birthdate,
		Biography:  raw.Biography,
		City:       raw.City,
	}, nil
}

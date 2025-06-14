package info

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"sn/internal/core"
)

const (
	LikeOperatorZeroOrMultipleSymbols = "%"
)

type (
	rawInfo struct {
		// fields are capitalized oly for sqlx correct parsing
		ID         uuid.UUID  `db:"account_id"`
		FirstName  *string    `db:"first_name"`
		SecondName *string    `db:"second_name"`
		Birthdate  *time.Time `db:"birth_date"`
		Biography  *string    `db:"biography"`
		City       *string    `db:"city"`
		CreatedAt  time.Time  `db:"created_at"`
		UpdatedAt  time.Time  `db:"updated_at"`
	}
)

func (s *infoStore) ReadInfo(ctx context.Context, accountID uuid.UUID) (core.PersonalInfoEntity, error) {
	sql := `SELECT
				account_id,
				first_name,
				second_name,
				birth_date,
				biography,
				city
			FROM personal_info
			WHERE account_id=$1`

	var raw rawInfo
	if err := s.client.DB.QueryRowxContext(ctx, sql, accountID).StructScan(&raw); err != nil {
		return core.PersonalInfoEntity{}, fmt.Errorf("fail to read an entry from `personal_info` table: %w", err)
	}

	return core.PersonalInfoEntity{
		UserID: raw.ID,
		PersonalInfo: core.PersonalInfo{
			FirstName:  raw.FirstName,
			SecondName: raw.SecondName,
			Birthdate:  raw.Birthdate,
			Biography:  raw.Biography,
			City:       raw.City,
		},
	}, nil
}

func (s *infoStore) GetInfoList(ctx context.Context, parameters core.SearchAccountsInfoParams) ([]core.PersonalInfoEntity, error) {
	sql := `SELECT
				account_id,
				first_name,
				second_name,
				birth_date,
				biography,
				city
			FROM personal_info
			WHERE first_name LIKE $1 AND second_name LIKE $2
			LIMIT $3`

	rows, err := s.client.DB.QueryxContext(
		ctx,
		sql,
		parameters.FirstName+LikeOperatorZeroOrMultipleSymbols,
		parameters.LastName+LikeOperatorZeroOrMultipleSymbols,
		parameters.Limit,
	)
	if err != nil {
		return []core.PersonalInfoEntity{}, fmt.Errorf("fail to search for entries in `personal_info` table: %w", err)
	}

	var raw rawInfo
	var entries []core.PersonalInfoEntity

	for rows.Next() {
		if err := rows.StructScan(&raw); err != nil {
			return []core.PersonalInfoEntity{}, fmt.Errorf("fail to convert `personal_info` table search result to `rawInfo` : %w", err)
		}

		entries = append(entries, core.PersonalInfoEntity{
			UserID: raw.ID,
			PersonalInfo: core.PersonalInfo{
				FirstName:  raw.FirstName,
				SecondName: raw.SecondName,
				Birthdate:  raw.Birthdate,
				Biography:  raw.Biography,
				City:       raw.City,
			},
		})
	}

	return entries, nil
}

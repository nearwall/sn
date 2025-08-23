package info

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"sn/internal/core"
)

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

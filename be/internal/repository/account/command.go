package account

import (
	"context"
	"fmt"
	"time"

	"sn/internal/core"
)

func (s *accountStore) Create(ctx context.Context, data core.AccountCreationData) error {
	hashedFeatures := createHashFeatures(data.Password)
	now := time.Now().UTC()
	sql := `INSERT INTO account (
				id,
				hashed_pwd,
				hash_features,
				pwd_updated_at,
				updated_at,
				created_at
			) VALUES ($1, $2, $3, $4, $5, $6)`

	if _, err := s.client.DB.ExecContext(ctx, sql, data.AccountID, data.Password.Hash, hashedFeatures, now, now, now); err != nil {
		return fmt.Errorf("fail to create an entry in `account` table: %w", err)
	}

	return nil
}

func createHashFeatures(password core.HashedPassword) int16 {
	var features int16
	switch password.Algorithm {
	case core.HashAlgoArgon2ID:
		features |= argonAlgo << hashAlgoStartBit
	case core.HashAlgoDebugBytesSum:
		features |= argonAlgo << hashAlgoStartBit
	}

	if password.Pepper != nil {
		features |= int16(password.Pepper.ID << hashAlgoStartBit)
	}

	return features
}

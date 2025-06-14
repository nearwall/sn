package account

import (
	"context"
	"errors"
	"fmt"

	"sn/internal/core"

	"github.com/google/uuid"
)

// core.AccountStore interface
func (s *accountStore) ReadPasswordInfo(ctx context.Context, accountID uuid.UUID) (core.PasswordInfo, error) {
	sql := `SELECT
				hashed_pwd,
				hash_features,
				pwd_updated_at
			FROM account
			WHERE id=$1`

	var raw passwordRawInfo
	if err := s.client.DB.QueryRowxContext(ctx, sql, accountID).StructScan(&raw); err != nil {
		return core.PasswordInfo{}, fmt.Errorf("fail to obtain `account` row: %w", err)
	}

	features, err := fromFeatures(raw.Features)
	if err != nil {
		return core.PasswordInfo{}, fmt.Errorf("fail to convert `account` row to `hashFeatures`: %w", err)
	}

	return core.PasswordInfo{
			Password: core.HashedPassword{
				Hash:      raw.Hash,
				Algorithm: features.algo,
				Pepper:    features.pepper,
			},
			UpdatedAt: raw.UpdateAt},
		nil
}

func fromFeatures(raw int16) (hashFeatures, error) {
	var features hashFeatures

	if (raw & pepperExistenceBit) != 0 {
		features.pepper.ID = uint8((raw & pepperIDMask) >> pepperIDStartBit)
	}

	switch (raw & hashAlgoMask) >> hashAlgoStartBit {
	case argonAlgo:
		features.algo = core.HashAlgoArgon2ID
		return features, nil
	case debugAlgo:
		features.algo = core.HashAlgoDebugBytesSum
		return features, nil
	default:
		return hashFeatures{}, errors.New("unknown hash algorithm: ")
	}
}

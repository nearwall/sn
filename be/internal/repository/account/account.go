package account

import (
	"context"
	"errors"
	"fmt"
	"time"

	"sn/internal/core"
	"sn/internal/infra/postgres"

	"github.com/google/uuid"
)

const (
	pepperExistenceBit = 0
	//  [1..3] (4 bits)
	hashAlgoStartBit = 1
	hashAlgoEndBit   = 3

	pepperIDStartBit = 4
	//  [4..11] (8 bits)
	pepperIDEndBit = 11
)

const (
	//  [1..3] (4 bits)
	hashAlgoMask = 0x000E
	//  [4..11] (8 bits)
	pepperIDMask = 0x0FF0
)

const (
	argonAlgo = 0
	// values [1..6] - reserved for future use
	debugAlgo = 7
)

type (
	passwordRawInfo struct {
		// fields are capitalized oly for sqlx correct parsing
		Hash     string    `db:"hashed_pwd"`
		Features int16     `db:"hash_features"`
		UpdateAt time.Time `db:"pwd_updated_at"`
	}

	hashFeatures struct {
		algo   core.PwdHashAlgorithm
		pepper *core.HashPepper
	}
)

type accountStore struct {
	client *postgres.Client
}

func NewAccountStore(postgresClient *postgres.Client) core.AccountStore {
	return &accountStore{
		client: postgresClient,
	}
}

// core.AccountStore interface
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

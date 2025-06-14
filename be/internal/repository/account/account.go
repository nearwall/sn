package account

import (
	"time"

	"sn/internal/core"
	"sn/internal/infra/postgres"
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

package core

import (
	"context"
)

type PwdHashAlgorithm uint8

const (
	HashAlgoArgon2ID PwdHashAlgorithm = iota

	HashAlgoDebugBytesSum = 255
)

type HashedPassword struct {
	Hash      string
	Algorithm PwdHashAlgorithm
	Pepper    HashPepper
}

type HashPepper struct {
	ID   uint8
	Used bool
}

type PasswordService interface {
	Hash(ctx context.Context, password string) (HashedPassword, error)
	Verify(ctx context.Context, password string, hashedPassword HashedPassword) (bool, error)
}

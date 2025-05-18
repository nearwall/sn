package core

import (
	"context"
)

type PwdHashAlgorithm uint8

const (
	Argon2ID PwdHashAlgorithm = iota

	DebugBytesSum = 255
)

type PasswordService interface {
	Hash(ctx context.Context, password string) (string, error)
	Verify(ctx context.Context, password string, hashedPassword string) (bool, error)
}

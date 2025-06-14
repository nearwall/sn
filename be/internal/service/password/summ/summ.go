package summ

import (
	"context"
	"errors"
	"strconv"

	"sn/internal/core"
)

type dbgPasswordService struct{}

func NewService() core.PasswordService {
	return &dbgPasswordService{}
}

// core.PasswordService interface
func (*dbgPasswordService) Hash(_ context.Context, password string) (core.HashedPassword, error) {
	return core.HashedPassword{Hash: strconv.Itoa(sumBytes(password)), Algorithm: core.HashAlgoDebugBytesSum}, nil
}

// core.PasswordService interface
func (*dbgPasswordService) Verify(_ context.Context, password string, hashedPassword core.HashedPassword) (bool, error) {
	if hashedPassword.Algorithm != core.HashAlgoDebugBytesSum {
		return false, errors.New("unsupported hash algorithm(only Debug Bytes Sum is available)")
	}
	hash := strconv.Itoa(sumBytes(password))

	return hash == hashedPassword.Hash, nil
}

func sumBytes(s string) int {
	sum := 0
	for i := range len(s) {
		sum += int(s[i])
	}
	return sum
}

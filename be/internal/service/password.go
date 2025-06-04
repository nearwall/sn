package service

import (
	"context"
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/crypto/argon2"

	"sn/internal/core"
)

type PasswordServiceConfig struct {
	HashPepper    string
	HashAlgorithm core.PwdHashAlgorithm
}

func NewPasswordService(config PasswordServiceConfig) core.PasswordService {
	switch config.HashAlgorithm {
	case core.HashAlgoArgon2ID:
		return &argon2PwdService{
			pepper: make([]byte, 0),
			Argon2IDConfig: Argon2IDConfig{
				Memory:      64 * 1024,
				Iterations:  4,
				Parallelism: 2,
				SaltLength:  16,
				KeyLength:   32,
			},
		}
	default:
		return &dbgPasswordService{}
	}
}

type (
	argon2PwdService struct {
		// ToDo: use pepper
		pepper         []byte
		Argon2IDConfig Argon2IDConfig
	}

	Argon2IDConfig struct {
		Memory      uint32
		Iterations  uint32
		Parallelism uint8
		SaltLength  uint32
		KeyLength   uint32
	}
)

var (
	ErrInvalidHash         = errors.New("the encoded hash is not in the correct format")
	ErrIncompatibleVersion = errors.New("incompatible version of argon2")
)

// core.PasswordService interface
func (p *argon2PwdService) Hash(ctx context.Context, password string) (core.HashedPassword, error) {
	salt, err := genSalt(p.Argon2IDConfig.SaltLength)
	if err != nil {
		return core.HashedPassword{}, err
	}

	hash := argon2.IDKey([]byte(password), salt, p.Argon2IDConfig.Iterations, p.Argon2IDConfig.Memory, p.Argon2IDConfig.Parallelism, p.Argon2IDConfig.KeyLength)

	// Base64 encode the salt and hashed password.
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	// Return a string using the standard encoded hash representation.
	encodedHash := fmt.Sprintf(
		"$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version,
		p.Argon2IDConfig.Memory,
		p.Argon2IDConfig.Iterations,
		p.Argon2IDConfig.Parallelism,
		b64Salt,
		b64Hash)

	return core.HashedPassword{Hash: encodedHash, Algorithm: core.HashAlgoArgon2ID}, nil
}

// core.PasswordService interface
func (_ *argon2PwdService) Verify(ctx context.Context, password string, hashedPassword core.HashedPassword) (bool, error) {
	if hashedPassword.Algorithm != core.HashAlgoArgon2ID {
		return false, fmt.Errorf("unsupported hash algorithm(only Argon2ID is available)")
	}
	hashParams, salt, hash, err := decodeHash(hashedPassword.Hash)
	if err != nil {
		return false, err
	}

	pwdHash := argon2.IDKey([]byte(password), salt, hashParams.Iterations, hashParams.Memory, hashParams.Parallelism, hashParams.KeyLength)

	// the subtle.ConstantTimeCompare() used to prevent timing attacks.
	if subtle.ConstantTimeCompare(hash, pwdHash) == 1 {
		return true, nil
	}

	return false, nil
}

func genSalt(bytesCnt uint32) ([]byte, error) {
	b := make([]byte, bytesCnt)
	if _, err := rand.Read(b); err != nil {
		return nil, err
	}

	return b, nil
}

func decodeHash(encodedHash string) (p Argon2IDConfig, salt, hash []byte, err error) {
	vals := strings.Split(encodedHash, "$")
	if len(vals) != 6 {
		return Argon2IDConfig{}, nil, nil, ErrInvalidHash
	}

	var version int
	_, err = fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		return Argon2IDConfig{}, nil, nil, err
	}
	if version != argon2.Version {
		return Argon2IDConfig{}, nil, nil, ErrIncompatibleVersion
	}

	p = Argon2IDConfig{}
	_, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &p.Memory, &p.Iterations, &p.Parallelism)
	if err != nil {
		return Argon2IDConfig{}, nil, nil, err
	}

	salt, err = base64.RawStdEncoding.Strict().DecodeString(vals[4])
	if err != nil {
		return Argon2IDConfig{}, nil, nil, err
	}
	p.SaltLength = uint32(len(salt))

	hash, err = base64.RawStdEncoding.Strict().DecodeString(vals[5])
	if err != nil {
		return Argon2IDConfig{}, nil, nil, err
	}
	p.KeyLength = uint32(len(hash))

	return p, salt, hash, nil
}

type dbgPasswordService struct{}

// core.PasswordService interface
func (_ *dbgPasswordService) Hash(ctx context.Context, password string) (core.HashedPassword, error) {
	return core.HashedPassword{Hash: strconv.Itoa(sumBytes(password)), Algorithm: core.HashAlgoDebugBytesSum}, nil
}

// core.PasswordService interface
func (_ *dbgPasswordService) Verify(ctx context.Context, password string, hashedPassword core.HashedPassword) (bool, error) {
	if hashedPassword.Algorithm != core.HashAlgoArgon2ID {
		return false, fmt.Errorf("unsupported hash algorithm(only Argon2ID is available)")
	}
	hash := strconv.Itoa(sumBytes(password))

	return hash == hashedPassword.Hash, nil
}

func sumBytes(s string) int {
	sum := 0
	for i := 0; i < len(s); i++ {
		sum += int(s[i])
	}
	return sum
}

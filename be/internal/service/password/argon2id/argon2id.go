package argon2id

import (
	"context"
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"

	"sn/internal/core"
)

type (
	argon2PwdService struct {
		// ToDo: use pepper
		pepper []byte
		config Config
	}

	Config struct {
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

func NewService(config Config, pepper []byte) core.PasswordService {
	return &argon2PwdService{
		pepper: pepper,
		config: config,
	}
}

// core.PasswordService interface
func (p *argon2PwdService) Hash(ctx context.Context, password string) (core.HashedPassword, error) {
	salt, err := genSalt(p.config.SaltLength)
	if err != nil {
		return core.HashedPassword{}, fmt.Errorf("fail to generate salt for hash: %w", err)
	}

	hash := argon2.IDKey(
		[]byte(password),
		salt,
		p.config.Iterations,
		p.config.Memory,
		p.config.Parallelism,
		p.config.KeyLength)

	// Base64 encode the salt and hashed password.
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	// Return a string using the standard encoded hash representation.
	encodedHash := fmt.Sprintf(
		"$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version,
		p.config.Memory,
		p.config.Iterations,
		p.config.Parallelism,
		b64Salt,
		b64Hash)

	return core.HashedPassword{Hash: encodedHash, Algorithm: core.HashAlgoArgon2ID}, nil
}

// core.PasswordService interface
func (*argon2PwdService) Verify(ctx context.Context, password string, hashedPassword core.HashedPassword) (bool, error) {
	if hashedPassword.Algorithm != core.HashAlgoArgon2ID {
		return false, errors.New("unsupported hash algorithm(only Argon2ID is available)")
	}
	hashParams, salt, hash, err := decodeHash(hashedPassword.Hash)
	if err != nil {
		return false, fmt.Errorf("fail to decode hash string: %w", err)
	}

	pwdHash := argon2.IDKey(
		[]byte(password),
		salt,
		hashParams.Iterations,
		hashParams.Memory,
		hashParams.Parallelism,
		hashParams.KeyLength)

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

func decodeHash(encodedHash string) (config Config, salt, hash []byte, err error) {
	vals := strings.Split(encodedHash, "$")
	if len(vals) != 6 {
		return Config{}, nil, nil, ErrInvalidHash
	}

	var version int
	_, err = fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		return Config{}, nil, nil, err
	}
	if version != argon2.Version {
		return Config{}, nil, nil, ErrIncompatibleVersion
	}

	config = Config{}
	_, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &config.Memory, &config.Iterations, &config.Parallelism)
	if err != nil {
		return Config{}, nil, nil, err
	}

	salt, err = base64.RawStdEncoding.Strict().DecodeString(vals[4])
	if err != nil {
		return Config{}, nil, nil, err
	}
	config.SaltLength = uint32(len(salt))

	hash, err = base64.RawStdEncoding.Strict().DecodeString(vals[5])
	if err != nil {
		return Config{}, nil, nil, err
	}
	config.KeyLength = uint32(len(hash))

	return
}

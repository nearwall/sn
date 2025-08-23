package password

import (
	"sn/internal/core"
)

type Config struct {
	HashPepper    *string
	HashAlgorithm core.PwdHashAlgorithm
}

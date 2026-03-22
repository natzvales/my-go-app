package auth

import (
	"crypto/rand"
	"encoding/base64"

	"github.com/natz/go-lib-app/internal/config"
	"golang.org/x/crypto/argon2"
)

// const (
// 	memory      = 64 * 1024
// 	iterations  = 3
// 	parallelism = 2
// 	keyLength   = 32
// 	saltLength  = 16
// )

func HashPassword(password string) (string, error) {

	cfg := config.LoadConfig()
	salt := make([]byte, cfg.SALTLENGTH)

	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey(
		[]byte(password),
		salt,
		uint32(cfg.ITERATIONS),
		uint32(cfg.MEMORY),
		uint8(cfg.PARALLELISM),
		uint32(cfg.KEYLENGTH),
	)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	return b64Salt + "." + b64Hash, nil
}

func VerifyPassword(password, encoded string) bool {

	cfg := config.LoadConfig()
	parts := split(encoded)
	if len(parts) != 2 {
		return false
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[0])
	if err != nil {
		return false
	}
	hash, err := base64.RawStdEncoding.DecodeString(parts[1])
	if err != nil {
		return false
	}

	newHash := argon2.IDKey(
		[]byte(password),
		salt,
		uint32(cfg.ITERATIONS),
		uint32(cfg.MEMORY),
		uint8(cfg.PARALLELISM),
		uint32(cfg.KEYLENGTH),
	)

	return base64.RawStdEncoding.EncodeToString(newHash) ==
		base64.RawStdEncoding.EncodeToString(hash)
}

func split(s string) []string {
	for i := 0; i < len(s); i++ {
		if s[i] == '.' {
			return []string{s[:i], s[i+1:]}
		}
	}
	return nil
}

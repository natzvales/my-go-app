package auth

import (
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/crypto/argon2"
)

const (
	memory      = 64 * 1024
	iterations  = 3
	parallelism = 2
	keyLength   = 32
	saltLength  = 16
)

func HashPassword(password string) (string, error) {

	salt := make([]byte, saltLength)

	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey(
		[]byte(password),
		salt,
		iterations,
		memory,
		parallelism,
		keyLength,
	)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	return b64Salt + "." + b64Hash, nil
}

func VerifyPassword(password, encoded string) bool {

	parts := split(encoded)

	salt, _ := base64.RawStdEncoding.DecodeString(parts[0])
	hash, _ := base64.RawStdEncoding.DecodeString(parts[1])

	newHash := argon2.IDKey(
		[]byte(password),
		salt,
		iterations,
		memory,
		parallelism,
		keyLength,
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

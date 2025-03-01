package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"strings"

	"golang.org/x/crypto/pbkdf2"
)

const (
	saltSize = 16
	iter     = 10000
	keyLen   = 32
)

func SaltPassword(password string) (string, error) {
	salt := make([]byte, saltSize)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	hash := pbkdf2.Key([]byte(password), salt, iter, keyLen, sha256.New)
	saltedHash := base64.StdEncoding.EncodeToString(hash) + ":" + base64.StdEncoding.EncodeToString(salt)
	return saltedHash, nil
}

func IsSaltMatched(password, saltedHash string) bool {
	parts := strings.Split(saltedHash, ":")
	if len(parts) != 2 {
		return false
	}

	hash, err := base64.StdEncoding.DecodeString(parts[0])
	if err != nil {
		return false
	}

	salt, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		return false
	}

	newHash := pbkdf2.Key([]byte(password), salt, iter, keyLen, sha256.New)
	return subtle.ConstantTimeCompare(hash, newHash) == 1
}

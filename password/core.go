package password

import (
	"crypto/rand"
	"crypto/sha512"
	"io"
)

const (
	ByteLenSalt = 16
)

func GenerateSalt() ([]byte, error) {
	salt := make([]byte, ByteLenSalt)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		return nil, err
	}
	return salt, nil
}

func Hash(password string, salt []byte) []byte {
	hasher := sha512.New()
	hasher.Write([]byte(password))
	hasher.Write(salt)
	return hasher.Sum(nil)
}

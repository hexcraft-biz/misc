package basic

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io"
	mrand "math/rand"
	"net/url"
	"os"
	"reflect"
	"time"
)

func IsSlice(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Slice
}

func UrlStandardize(source string) (string, error) {
	if u, err := url.Parse(source); err != nil {
		return "", err
	} else {

		query := u.Query().Encode()
		if query != "" {
			query = "?" + query
		}

		return (u.Scheme + "://" + u.Host + u.EscapedPath() + query), nil
	}
}

// ================================================================
//
// ================================================================
const (
	DefCharsetNumber uint8 = 1 << iota
	DefCharsetLowercase
	DefCharsetUppercase
	DefCharsetSpecialChars
	DefCharsetNotSet uint8 = 0x00
	DefCharsetAll    uint8 = DefCharsetNumber | DefCharsetLowercase | DefCharsetUppercase | DefCharsetSpecialChars

	DefNumbers      string = "0123456789"
	DefLowercases   string = "abcdefghijklmnopqrstuvwxyz"
	DefUppercases   string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	DefSpecialChars string = "~`!@#$%^&*()-_=+[{]}||;:',<.>/?"
)

func GenStringWithCharset(length int, charset uint8) string {
	chars := ""

	if (charset & DefCharsetNumber) > 0 {
		chars += DefNumbers
	}
	if (charset & DefCharsetLowercase) > 0 {
		chars += DefLowercases
	}
	if (charset & DefCharsetUppercase) > 0 {
		chars += DefUppercases
	}
	if (charset & DefCharsetSpecialChars) > 0 {
		chars += DefSpecialChars
	}

	charlen := len(chars)

	seed := mrand.New(mrand.NewSource(time.Now().UnixNano()))
	s := make([]byte, length)
	for i := range s {
		s[i] = chars[seed.Intn(charlen)]
	}

	return string(s)
}

func GenerateSalt(size int) ([]byte, error) {
	salt := make([]byte, size)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		return nil, err
	}
	return salt, nil
}

func HashSha512(password string, salt []byte) []byte {
	hasher := sha512.New()
	hasher.Write([]byte(password))
	hasher.Write(salt)
	return hasher.Sum(nil)
}

// ================================================================
//
// ================================================================
func CompareChecksum(chksum string, fp string) (bool, error) {
	if databytes, err := os.ReadFile(fp); err != nil {
		return false, err
	} else if chksum == fmt.Sprintf("%x", sha256.Sum256(databytes)) {
		return true, nil
	} else {
		return false, nil
	}
}

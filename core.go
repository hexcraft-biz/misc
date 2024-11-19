package misc

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io"
	"math/big"
	mrand "math/rand"
	"net/url"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/hexcraft-biz/xuuid"
)

func IsSlice(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Slice
}

func UrlStandardize(source string) (string, error) {
	if u, err := url.Parse(source); err != nil {
		return "", err
	} else {
		u.RawQuery = u.Query().Encode()
		return u.String(), nil
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

func GenerateSha512Hmac(password string, salt []byte) []byte {
	h := hmac.New(sha512.New, salt)
	h.Write([]byte(password))
	return h.Sum(nil)
}

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

// ----------------------------------------------------------------
const base62Chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func base62Encode(num *big.Int) string {
	if num.Sign() == 0 {
		return "0"
	}

	var encoded string
	base := big.NewInt(62)
	zero := big.NewInt(0)

	for num.Cmp(zero) > 0 {
		mod := new(big.Int)
		num.DivMod(num, base, mod)
		encoded = string(base62Chars[mod.Int64()]) + encoded
	}

	return encoded
}

func uuidToBase62(id xuuid.UUID) string {
	noDashUuidStr := strings.ReplaceAll(id.String(), "-", "")
	num := new(big.Int)
	num.SetString(noDashUuidStr, 16)
	return base62Encode(num)
}

package crypto

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

// Md5 .
func Md5(str string) string {
	hash := md5.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}

// Sha512 .
func Sha512(str string) string {
	hash := sha512.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}

// Sha256 .
func Sha256(str string) string {
	hash := sha256.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}

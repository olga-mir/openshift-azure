package random

import (
	"crypto/rand"
	"io"
	"math/big"
)

// String returns a random string of length n comprised of bytes in letterBytes
func String(letterBytes string, n int) (string, error) {
	b := make([]byte, n)
	for i := range b {
		o, err := rand.Int(rand.Reader, big.NewInt(int64(len(letterBytes))))
		if err != nil {
			return "", err
		}
		b[i] = letterBytes[o.Int64()]
	}

	return string(b), nil
}

// AlphanumericString returns a random string of length n comprised of
// [A-Za-z0-9]
func AlphanumericString(n int) (string, error) {
	return String("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", n)
}

// LowerCaseAlphanumericString returns a random string of length n comprised of
// [a-z0-9]
func LowerCaseAlphanumericString(n int) (string, error) {
	return String("abcdefghijklmnopqrstuvwxyz0123456789", n)
}

// Bytes returns a random byte slice of legth n
func Bytes(n int) ([]byte, error) {
	b := make([]byte, n)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return nil, err
	}
	return b, nil
}

// StorageAccountName returns a random string suitable for use as an Azure
// storage account name
func StorageAccountName() (string, error) {
	return LowerCaseAlphanumericString(24)
}

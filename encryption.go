package exgo

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

// hmac sha256
func HmacSHA256(data, key []byte) (string, error) {
	hash := hmac.New(sha256.New, key)
	if _, err := hash.Write(data); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(hash.Sum(nil)), nil
}

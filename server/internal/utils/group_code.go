package utils

import (
	"crypto/rand"
	"math/big"
)

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateGroupCode(length int) (string, error) {
	groupCode := make([]byte, length)
	for i := range groupCode {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		groupCode[i] = charset[index.Int64()]
	}
	return string(groupCode), nil
}
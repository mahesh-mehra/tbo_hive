package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func GenerateOTP() (string, error) {
	// Generate a random number between 0 and 9999
	n, err := rand.Int(rand.Reader, big.NewInt(10000))
	if err != nil {
		return "", err
	}
	// Format as 4 digits (pads with leading zeros if needed, e.g., "0042")
	return fmt.Sprintf("%04d", n), nil
}

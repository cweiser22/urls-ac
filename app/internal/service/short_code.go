package service

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
)

const base62Alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

type ShortCodeService struct {
}

// base62Encode encodes a byte slice into a base62 string
func base62Encode(data []byte) string {
	num := new(big.Int).SetBytes(data)
	base := big.NewInt(62)
	var encoded []byte

	for num.Sign() > 0 {
		mod := new(big.Int)
		num.DivMod(num, base, mod)
		encoded = append([]byte{base62Alphabet[mod.Int64()]}, encoded...)
	}
	return string(encoded)
}

func (s *ShortCodeService) GenerateShortcode(longURL string, length int) string {
	// Generate a 4-byte random salt
	salt := make([]byte, 4)
	if _, err := rand.Read(salt); err != nil {
		panic("failed to generate random salt: " + err.Error())
	}

	// Create a hash input: URL + salt + length to avoid accidental collisions
	input := fmt.Sprintf("%s|%x|%d", longURL, salt, length)

	// Hash using SHA-256
	hash := sha256.Sum256([]byte(input))

	// Encode to Base62
	base62 := base62Encode(hash[:])

	// Return the first N characters
	return base62[:length]
}

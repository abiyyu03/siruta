package helper

import (
	cryptoRand "crypto/rand"
	"encoding/hex"
	"io"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

const letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// RandomString generates a random string of a specified length.
func RandomString(length int) string {
	b := make([]byte, length)
	seed := rand.NewSource(time.Now().UnixNano()) // Seed for randomness
	r := rand.New(seed)

	for i := range b {
		b[i] = letterBytes[r.Intn(len(letterBytes))]
	}
	return string(b)
}

func ConvertPascalCaseToSnakeCase(input string) string {
	// Use regex to insert underscores before capital letters
	regex := regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := regex.ReplaceAllString(input, "${1}_${2}")

	// Convert to lowercase
	return strings.ToLower(snake)
}

func GenerateSecureToken(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := io.ReadFull(cryptoRand.Reader, bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

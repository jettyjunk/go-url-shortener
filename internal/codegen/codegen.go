package codegen

import (
	"crypto/rand"
	"errors"
	"fmt"
)

const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Generate(lenght int) (string, error) {
	if lenght <= 0 {
		return "", errors.New("length must be positive")
	}

	byt := make([]byte, lenght)

	if _, err := rand.Read(byt); err != nil {
		return "", fmt.Errorf("generate code: %w", err)
	}

	for i, v := range byt {
		byt[i] = alphabet[int(v)%len(alphabet)]
	}

	return string(byt), nil
}

package util

import (
	"errors"
	"fmt"
	"os"
)

func GetEnv(key string, fallback string) (string, error) {
	var value string
	if value = os.Getenv(key); value == "" {
		if fallback == "" {
			return "", errors.New(fmt.Sprintf("OS environment variable: %s not set and no fallback value set", key))
		}
		return fallback, nil
	}
	return value, nil
}

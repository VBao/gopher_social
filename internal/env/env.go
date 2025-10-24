package env

import (
	"os"
	"strconv"
)

func GetString(key, fallback string) string {
	val, isFound := os.LookupEnv(key)

	if isFound {
		return val
	}

	return fallback
}

func GetInt(key string, fallback int) int {
	valStr := GetString(key, strconv.Itoa(fallback))

	val, err := strconv.Atoi(valStr)
	if err != nil {
		return fallback
	}

	return val
}

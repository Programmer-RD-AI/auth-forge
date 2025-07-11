package authforge

import (
	"os"
	"strconv"
)

func getEnv(key string, defaultVal any) any {
	rawEnv, ok := os.LookupEnv(key)
	if !ok{
		return defaultVal
	}
	if intEnv, err := strconv.Atoi(rawEnv); err == nil {
		return intEnv
	}
	return rawEnv;
}


package util

import (
	"os"
	"strconv"
)

func GetEnv[T string | int](key string, defaultVal T) (T, error) {
	rawEnv, ok := os.LookupEnv(key)
	if !ok {
		return defaultVal, nil
	}
	var ret any
	switch any(defaultVal).(type) {
	case string:
		ret = rawEnv
	case int:
		intVal, err := strconv.Atoi(rawEnv)
		if err != nil {
			return defaultVal, err
		}
		ret = intVal

	}
	return ret.(T), nil
}

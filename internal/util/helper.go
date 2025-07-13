package util

import (
	"encoding/json"
)

func MarshalBinary[T any](data T) ([]byte, error) {
	return json.Marshal(data)
}

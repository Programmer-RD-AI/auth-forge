package authforge

import (
	"fmt"
)

type DbHealthCheckFail struct {
	message  string
	provider string
}

func (e *DbHealthCheckFail) Error() string {
	return fmt.Sprintf("[%s] %s", e.provider, e.message)
}
func NewDbHealthCheckFail(provider string, message string) error {
	return &DbHealthCheckFail{provider: provider, message: message}
}

type KeyDoesNotExistError struct {
	key string
}

func (k *KeyDoesNotExistError) Error() string {
	return fmt.Sprintf("Key '%s' already exists in Redis", k.key)
}

func NewKeyDoesNotExistError(key string) error {
	return &KeyDoesNotExistError{key: key}
}

type UserIdRequiredForSession struct {
	providedSessionValue any
}

func (r *UserIdRequiredForSession) Error() string {
	return fmt.Sprintf("User Id is required for Session Value Creation (Provided Session Value: %v", r.providedSessionValue)
}

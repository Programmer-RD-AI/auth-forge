package errors

import "fmt"

type DbHealthCheckFail struct {
	message  string
	provider string
}

func (e *DbHealthCheckFail) Error() string {
	return fmt.Sprintf("[%s] %s", e.provider, e.message)
}
func NewDbHealthCheckFail(provider, message string) error {
	return &DbHealthCheckFail{provider: provider, message: message}
}

type KeyDoesNotExistError struct {
	key string
}

func (k *KeyDoesNotExistError) Error() string {
	return fmt.Sprintf("Key '%s' already exists in Redis", k.key)
}

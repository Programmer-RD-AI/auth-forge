package errors

import "fmt"

type dbHealthCheckFail struct {
	message  string
	provider string
}

func (e *dbHealthCheckFail) Error() string {
	return fmt.Sprintf("[%s] %s", e.provider, e.message);
}


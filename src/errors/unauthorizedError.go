package errors

import (
	"fmt"
	"net/http"
)

func NewUnauthorizedError(response *http.Response) *UnauthorizedError {
	return &UnauthorizedError{response}
}

type UnauthorizedError struct {
	response *http.Response
}

func (this *UnauthorizedError) Error() string {
	return fmt.Sprintf("%s", this.response.Status)
}

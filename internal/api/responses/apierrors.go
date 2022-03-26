package responses

import (
	"fmt"
	"log"
)

type CauseList []interface{}

type (
	apiError struct {
		Message string    `json:"message"`
		Cause   CauseList `json:"cause"`
		Error   string    `json:"error"`
	}

	ApiError interface {
		GetMessage() string
		GetCause() CauseList
		GetError() string
	}
)

// NewApiError returns a new instance of ApiError
func NewApiError(message string, cause CauseList, error string) ApiError {
	return &apiError{message, cause, error}
}

func NewNotFound(resource string) ApiError {
	args := make([]interface{}, 1)
	args[0] = resource
	return &apiError{fmt.Sprintf("Resource %s not found", resource), args, "NotFound"}
}

func NewBadRequest(err string) ApiError {
	args := make([]interface{}, 1)
	args[0] = err
	return &apiError{fmt.Sprintf("Malformed body: %s ", err), args, "BadRequest"}
}

func NewInternalServerError(err string) ApiError {
	log.Printf("[InternalServerError] %s", err)
	args := make([]interface{}, 1)
	args[0] = err
	return &apiError{fmt.Sprintf("Ups, something broke inside :S"), args, "InternalServerError"}
}

func (ae apiError) GetMessage() string {
	return ae.Message
}

func (ae apiError) GetCause() CauseList {
	return ae.Cause
}

func (ae apiError) GetError() string {
	return ae.Error
}

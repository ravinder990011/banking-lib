package errs

import "net/http"

type AppErr struct {
	Code    int    `json:"-"`
	Message string `json:"message"`
}

func NewNotFoundError(message string) *AppErr {
	return &AppErr{
		Message: message,
		Code:    http.StatusNotFound,
	}
}
func NewUnexpectedError(message string) *AppErr {
	return &AppErr{
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}

func NewValiationError(message string) *AppErr {
	return &AppErr{
		Message: message,
		Code:    http.StatusUnprocessableEntity,
	}
}

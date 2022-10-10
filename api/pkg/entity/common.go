package entity

import "net/http"

type AppError struct {
	HTTPCode     int
	InternalCode int
	Error        error
	Message      string
}

func NewAppError(internalCode int) *AppError {
	return &AppError{
		InternalCode: internalCode,
		// Set defaults
		HTTPCode: http.StatusInternalServerError,
		Message:  "An unexpected error occurred. Please try again",
	}
}

func (a *AppError) SetError(err error) *AppError {
	a.Error = err
	return a
}

func (a *AppError) SetMessage(msg string) *AppError {
	a.Message = msg
	return a
}

func (a *AppError) SetHTTPCode(httpCode int) *AppError {
	a.HTTPCode = httpCode
	return a
}

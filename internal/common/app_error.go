package common

import "net/http"

type ErrorCode string

const (
	BadRequest            ErrorCode = "BAD_REQUEST"
	ValidationError       ErrorCode = "VALIDATION_ERROR"
	AuthRequired          ErrorCode = "AUTH_REQUIRED"
	AuthInvalidCredential ErrorCode = "AUTH_INVALID_CREDENTIALS"
	NotFound              ErrorCode = "NOT_FOUND"
	Forbidden             ErrorCode = "FORBIDDEN"
	ServerError           ErrorCode = "SERVER_ERROR"
	DatabaseError         ErrorCode = "DATABASE_ERROR"
	TooManyRequests       ErrorCode = "TOO_MANY_REQUESTS"
)

type AppError struct {
	StatusCode int
	Code       ErrorCode
	Message    string
	Details    any
}

func (error *AppError) Error() string {
	return error.Message
}

func NewAppError(statusCode int, code ErrorCode, message string, details any) *AppError {
	return &AppError{
		StatusCode: statusCode,
		Code:       code,
		Message:    message,
		Details:    details,
	}
}

func BadRequestError(details any) *AppError {
	return NewAppError(http.StatusBadRequest, BadRequest, "Invalid request body", details)
}

func NotFoundError(message string) *AppError {
	return NewAppError(http.StatusNotFound, NotFound, message, nil)
}

func InternalServerError() *AppError {
	return NewAppError(http.StatusInternalServerError, ServerError, "An unexpected error occurred", nil)
}

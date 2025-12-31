package common

import (
	"encoding/json"
	"errors"
	"io"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

func CheckTypeError(typeError *json.UnmarshalTypeError, errorMap map[string]string) map[string]string {
	field := strings.ToLower(typeError.Field)
	kind := typeError.Type.Kind()

	// Inisialisasi pesan default
	defaultMessage := field + " has invalid type"

	// Periksa kondisi menggunakan if terpisah
	if kind == reflect.Float32 || kind == reflect.Float64 {
		errorMap[field] = field + " must be a number"
	}

	if kind == reflect.Int || kind == reflect.Int64 {
		errorMap[field] = field + " must be an integer"
	}

	if kind == reflect.String {
		errorMap[field] = field + " must be a string"
	}

	if kind == reflect.Bool {
		errorMap[field] = field + " must be a boolean"
	}

	if errorMap[field] == "" {
		errorMap[field] = defaultMessage
	}

	return errorMap
}

func ErrorValidation(err error) map[string]string {
	errorsMap := make(map[string]string)

	// Body kosong
	if errors.Is(err, io.EOF) {
		errorsMap["request"] = "Request body is required"
		return errorsMap
	}

	// Salah tipe data (json unmarshal)
	var typeErr *json.UnmarshalTypeError
	if errors.As(err, &typeErr) {
		return CheckTypeError(typeErr, errorsMap)
	}

	// Validation errors (binding tag)
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range validationErrors {
			field := strings.ToLower(fieldErr.Field())

			switch fieldErr.Tag() {
			case "uuid":
				errorsMap[field] = field + " must be a valid UUID"
			case "required":
				errorsMap[field] = field + " is required"
			case "min":
				errorsMap[field] = field + " must be at least " + fieldErr.Param() + " characters"
			default:
				errorsMap[field] = field + " is invalid"
			}
		}
		return errorsMap
	}

	// Fallback
	errorsMap["request"] = "Invalid request payload"
	return errorsMap
}

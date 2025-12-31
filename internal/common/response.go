package common

type ErrorSchema struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
	Details any       `json:"details,omitempty"`
}

type Response[T any] struct {
	Success bool         `json:"success"`
	Data    *T           `json:"data,omitempty"`
	Errors  *ErrorSchema `json:"errors,omitempty"`
}

func SuccessResponse[T any](data T) Response[T] {
	return Response[T]{
		Success: true,
		Data:    &data,
		Errors:  nil,
	}
}

func ErrorResponse[T any](errors ErrorSchema) Response[T] {
	return Response[T]{
		Success: false,
		Data:    nil,
		Errors:  &errors,
	}
}

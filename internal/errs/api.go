package errs

import (
	"net/http"
)

type ApiError struct {
	StatusCode  uint
	Description string
}

func (a ApiError) Error() string {
	return a.Description
}

func (a ApiError) Is(err error) bool {
	apiError, ok := err.(ApiError)
	return ok && apiError.StatusCode == a.StatusCode
}

func NotFound(descripton string) ApiError {
	return ApiError{
		Description: descripton,
		StatusCode:  http.StatusNotFound,
	}
}

func UnprocessableEntity(descripton string) ApiError {
	return ApiError{
		Description: descripton,
		StatusCode:  http.StatusUnprocessableEntity,
	}
}

package validation

import (
	"fmt"
)

type ErrorDetail struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ErrorList []ErrorDetail

func RequiredError(field string) ErrorDetail {
	return ErrorDetail{
		Field:   field,
		Message: fmt.Sprintf("'%s' is required", field),
	}
}

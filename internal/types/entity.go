package types

import (
	"github.com/pmoura-dev/esr-service/internal/validation"
)

type Entity struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (e Entity) Validate() validation.ErrorList {
	errorList := validation.ErrorList{}

	if e.ID == "" {
		errorList = append(errorList, validation.RequiredError("id"))
	}

	if e.Name == "" {
		errorList = append(errorList, validation.RequiredError("name"))
	}

	return errorList
}

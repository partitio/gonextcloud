package types

import (
	"fmt"
	"strings"
)

type APIError struct {
	Code    int
	Message string
}

func ErrorFromMeta(meta Meta) *APIError {
	return &APIError{
		meta.Statuscode,
		meta.Message,
	}
}

func (e *APIError) Error() string {
	return fmt.Sprintf("%d : %s", e.Code, e.Message)
}

type UpdateError struct {
	Field string
	Error error
}

type UserUpdateError struct {
	Errors map[string]error
}

func (e *UserUpdateError) Error() string {
	var errors []string
	for k, e := range e.Errors {
		errors = append(errors, fmt.Sprintf("%s: %s", k, e.Error()))
	}
	return strings.Join(errors, ",")
}

func NewUpdateError(errors chan UpdateError) *UserUpdateError {
	empty := true
	var ue UserUpdateError
	for e := range errors {
		if ue.Errors == nil {
			empty = false
			ue.Errors = map[string]error{e.Field: e.Error}
		}
		ue.Errors[e.Field] = e.Error
	}
	if !empty {
		return &ue
	}
	return nil
}

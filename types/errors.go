package types

import (
	"fmt"
	"strings"
)

//APIError contains the returned error code and message from the Nextcloud's API
type APIError struct {
	Code    int
	Message string
}

//ErrorFromMeta return a types.APIError from the Response's types.Meta
func ErrorFromMeta(meta Meta) *APIError {
	return &APIError{
		meta.Statuscode,
		meta.Message,
	}
}

//Error return the types.APIError string
func (e *APIError) Error() string {
	return fmt.Sprintf("%d : %s", e.Code, e.Message)
}

//UpdateError contains the user's field and corresponding error
type UpdateError struct {
	Field string
	Error error
}

//UpdateError contains the errors resulting from a UserUpdate or a UserCreateFull call
type UserUpdateError struct {
	Errors map[string]error
}

func (e *UserUpdateError) Error() string {
	var errors []string
	for k, e := range e.Errors {
		errors = append(errors, fmt.Sprintf("%s: %v", k, e))
	}
	return strings.Join(errors, ",")
}

//NewUpdateError returns an UpdateError based on an UpdateError channel
func NewUpdateError(errors chan UpdateError) *UserUpdateError {
	var ue UserUpdateError
	for e := range errors {
		if ue.Errors == nil {
			ue.Errors = map[string]error{e.Field: e.Error}
		}
		ue.Errors[e.Field] = e.Error
	}
	if len(ue.Errors) > 0 {
		return &ue
	}
	return nil
}

package types

import "fmt"

type APIError struct {
	Code    int
	Message string
}

func ErrorFromMeta(meta Meta) APIError {
	return APIError{
		meta.Statuscode,
		meta.Message,
	}
}

func (e *APIError) Error() string {
	return fmt.Sprintf("%d : %s", e.Code, e.Message)
}

package utils

import "errors"

var (
	ErrDuplicateEmail    = errors.New("duplicate email")
	ErrDuplicateUsername = errors.New("duplicate username")
	ErrNotFound          = errors.New("record not found")
	ErrUnAuthorized      = errors.New("unauthorized")
	ErrInternal          = errors.New("internal error")
	ErrNotCreated        = errors.New("not created")
)

type Error struct {
	Errors map[string]interface{} `json:"errors"`
}

func NewError(err error) Error {
	e := Error{}
	e.Errors = make(map[string]interface{})
	switch v := err.(type) {
	default:
		e.Errors["body"] = v.Error()
	}

	return e
}

func AccessForbidden() Error {
	e := Error{}
	e.Errors = make(map[string]interface{})
	e.Errors["body"] = "access forbidden"
	return e
}

package users

import "errors"

var (
	ErrEmailNotUnique = errors.New(
		"email already in use",
	)
)

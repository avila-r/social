package posts

import "errors"

var (
	ErrInvalidSenderID = errors.New(
		"sender doesn't exist",
	)
)

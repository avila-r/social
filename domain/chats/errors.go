package chats

import (
	"github.com/avila-r/xgo/pkg/api"
)

var (
	ErrWebsocketRequired = &api.Error{
		// Upgrade required
		Code:    426,
		Message: "this endpoint requires a web socket connection",
	}

	ErrChatAlreadyExists = &api.Error{
		// Conflict
		Code:    409,
		Message: "chat already exists",
	}
)

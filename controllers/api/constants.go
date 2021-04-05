package apicontrollers

import (
	"errors"
)

const ContentType = "application/vnd.api+json; charset=utf-8"

var (
	ErrorInvalidUser        = errors.New("invalid user")
	ErrorMissingAccessToken = errors.New("missing access token")
	ErrorNotFoundUser       = errors.New("user not found")
)

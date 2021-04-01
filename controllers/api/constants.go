package apicontrollers

import (
	"errors"
)

const (
	ContentType     = "application/vnd.api+json; charset=utf-8"
	DefaultPageSize = 10
)

var (
	DefaultOrderBy = []string{"created_at desc"}

	ErrorInvalidUser           = errors.New("invalid user")
	ErrorMissingAccessToken    = errors.New("missing access token")
	ErrorNotFoundUser          = errors.New("user not found")
	ErrorRetrieveKeywordFailed = errors.New("there was a problem retrieving all keywords")
)

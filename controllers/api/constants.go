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
	ErrorInvalidPayloaderType  = errors.New("invalid payload type")
	ErrorMissingAccessToken    = errors.New("missing access token")
	ErrorNotFoundReport        = errors.New("report not found")
	ErrorNotFoundUser          = errors.New("user not found")
	ErrorGenerateReportFailed  = errors.New("there was a problem generating a report")
	ErrorRetrieveKeywordFailed = errors.New("there was a problem retrieving all keywords")
)

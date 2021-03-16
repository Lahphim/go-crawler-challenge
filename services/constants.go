package services

const ContentMinimumSize = 1
const ContentMaximumSize = 1000

var ValidationMessages = map[string]string{
	"InvalidLinkList":   "All Link list must be valid URL",
	"InvalidUrl":        "URL must be valid",
	"InvalidUser":       "User does not exist",
	"ExceedKeywordSize": "Acceptance keyword size from 1 to 1,000",
}

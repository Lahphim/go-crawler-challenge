package forms

const KeywordUploadContentTypeCSV = "text/csv"
const KeywordUploadMinimumSize = 1
const KeywordUploadMaximumSize = 1000

var ValidationMessages = map[string]string{
	"ConfirmPassword":    "Confirm password confirmation must match",
	"ExistingEmail":      "Email is already in use",
	"InvalidCredential":  "Your email or password is incorrect",
	"InvalidFileType":    "File type is not allowed",
	"InvalidKeywordSize": "Acceptance keyword size from 1 to 1,000",
	"OpenFile":           "File cannot be opened",
	"RequireFile":        "File cannot be empty",
	"ValidationFailed":   "Set validation error failed",
}

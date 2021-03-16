package forms

const ContentTypeCSV = "text/csv"
const ContentMinimumSize = 1
const ContentMaximumSize = 1000

var ValidationMessages = map[string]string{
	"ConfirmPassword":       "Confirm password confirmation must match",
	"ExistingEmail":         "Email is already in use",
	"InvalidCredential":     "Your email or password is incorrect",
	"InvalidUploadFileType": "File type is not allowed",
	"ExceedKeywordSize":     "Acceptance keyword size from 1 to 1,000",
	"OpenUploadFile":        "File cannot be opened",
	"RequireUploadFile":     "File cannot be empty",
}

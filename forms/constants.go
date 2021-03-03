package forms

var ValidationMessages = map[string]string{
	"ConfirmPassword":   "Confirm password confirmation must match",
	"ExistingEmail":     "Email is already in use",
	"InvalidCredential": "Your email or password is incorrect",
	"InvalidLinkList":   "All Link list must be valid URL",
	"InvalidUrl":        "URL must be valid",
	"InvalidUser":       "User does not exist",
}

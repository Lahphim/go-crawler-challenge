package models

import (
	"time"

	"github.com/beego/beego/v2/core/validation"
)

type Base struct {
	Id        int64
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
}

func init() {
	// Set default messages
	validation.SetDefaultMessage(map[string]string{
		"Required":     "can not be empty",
		"Min":          "minimum is %d",
		"Max":          "maximum is %d",
		"Range":        "range is %d to %d",
		"MinSize":      "minimum size is %d",
		"MaxSize":      "maximum size is %d",
		"Length":       "required length is %d",
		"Alpha":        "must be valid alpha characters",
		"Numeric":      "must be valid numeric characters",
		"AlphaNumeric": "must be valid alpha or numeric characters",
		"Match":        "must match %s",
		"NoMatch":      "must not match %s",
		"AlphaDash":    "must be valid alpha or numeric or dash(-_) characters",
		"Email":        "must be a valid email address",
		"IP":           "must be a valid ip address",
		"Base64":       "must be valid base64 characters",
		"Mobile":       "must be valid mobile number",
		"Tel":          "must be valid telephone number",
		"Phone":        "must be valid telephone or mobile phone number",
		"ZipCode":      "must be valid zipcode",
	})
}

package helpers

import (
	"crypto/md5"
	"encoding/hex"
	"html/template"
	"time"

	"github.com/justincampbell/timeago"
)

// HashEmail converts plain email string into hashed value using for render unique gravatar
// following this implementation here https://en.gravatar.com/site/implement/images/
func HashEmail(plainEmail string) string {
	byteEmail := md5.Sum([]byte(plainEmail))

	return hex.EncodeToString(byteEmail[:])
}

// ToTimeAgo converts timestamp to human readable in time ago format
func ToTimeAgo(timestamp time.Time) string {
	return timeago.FromTime(timestamp)
}

// ToTimeStamp converts timestamp to custom format
func ToTimeStamp(timestamp time.Time) string {
	const layout = "01/02/2006 | 3:04PM"

	return timestamp.Local().Format(layout)
}

// Unescape some raw HTML to displayable in screen
func Unescape(rawHtml string) template.HTML {
	return template.HTML(rawHtml)
}

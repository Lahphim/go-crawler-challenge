package models

import (
	"time"
)

// Report : Report model
type Report struct {
	Id      int64
	Keyword string
	Url     string
	RawHtml string

	LinkList map[string][]string

	CreatedAt time.Time
	UpdatedAt time.Time

	TotalAdsTop int
	TotalAds    int
	TotalNonAds int
	TotalLink   int
}

package models

import "time"

// Base : Default base struct for every model
type Base struct {
	Id        uint
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
}

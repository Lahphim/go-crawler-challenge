package models

import "time"

type Base struct {
	Id        int64     `orm:"auto"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
}

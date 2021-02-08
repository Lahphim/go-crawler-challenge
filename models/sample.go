package models

import (
	"github.com/beego/beego/v2/client/orm"
)

// Sample : Sample model
type Sample struct {
	Id          int64
	Title       string `orm:"size(128)"`
	Description string `orm:"type(longtext)"`
}

func init() {
	orm.RegisterModel(new(Sample))
}

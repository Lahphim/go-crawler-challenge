package models

import (
	"github.com/beego/beego/v2/client/orm"
)

// Page : Page model
type Page struct {
	Base

	Keyword *Keyword `orm:"rel(one)"`

	RawHtml string `orm:"type(text)"`
}

func init() {
	orm.RegisterModel(new(Page))
}

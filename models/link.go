package models

import (
	"github.com/beego/beego/v2/client/orm"
)

// Link : Link model
type Link struct {
	Base

	Keyword  *Keyword  `orm:"rel(fk)"`
	Position *Position `orm:"null;rel(fk);on_delete(set_null)"`

	Url string `orm:"type(text)"`
}

func init() {
	orm.RegisterModel(new(Link))
}

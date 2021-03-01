package models

import (
	"github.com/beego/beego/v2/client/orm"
)

// Keyword : Keyword model
type Keyword struct {
	Base

	User  *User   `orm:"rel(fk)"`
	Page  *Page   `orm:"reverse(one)"`
	Links []*Link `orm:"reverse(many)"`

	Keyword  string `orm:"size(128)"`
	VisitUrl string `orm:"size(128)"`
}

func init() {
	orm.RegisterModel(new(Keyword))
}

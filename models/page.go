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

// AddPage insert a new Page into database and returns last inserted Id on success.
func AddPage(page *Page) (id int64, err error) {
	ormer := orm.NewOrm()
	id, err = ormer.Insert(page)

	return id, err
}

package models

import (
	"github.com/beego/beego/v2/client/orm"
)

// Link : Link model
type Link struct {
	Base

	Keyword  *Keyword  `orm:"rel(fk)"`
	Position *Position `orm:"null;rel(fk);on_delete(set_null)"`

	Url string `orm:"size(128)"`
}

func init() {
	orm.RegisterModel(new(Link))
}

// AddLink insert a new Link into database and returns last inserted Id on success.
func AddLink(link *Link) (id int64, err error) {
	ormer := orm.NewOrm()
	id, err = ormer.Insert(link)

	return id, err
}

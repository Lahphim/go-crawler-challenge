package models

import (
	"github.com/beego/beego/v2/client/orm"
)

// Position : Position model
type Position struct {
	Base

	Name     string `orm:"size(128)"`
	Selector string `orm:"size(128)"`
	Category string `orm:"size(128)"`
}

func init() {
	orm.RegisterModel(new(Position))
}

// AddPosition insert a new Position into database and returns last inserted Id on success.
func AddPosition(position *Position) (id int64, err error) {
	ormer := orm.NewOrm()
	id, err = ormer.Insert(position)

	return id, err
}

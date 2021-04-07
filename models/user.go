package models

import (
	"github.com/beego/beego/v2/client/orm"
)

// User : User model
type User struct {
	Base

	Keywords []*Keyword `orm:"reverse(many)"`

	Email          string `orm:"unique;size(128)"`
	HashedPassword string `orm:"size(128)"`
}

func init() {
	orm.RegisterModel(new(User))
}

// AddUser insert a new User into database and returns last inserted Id on success.
func AddUser(user *User) (id int64, err error) {
	ormer := orm.NewOrm()
	id, err = ormer.Insert(user)

	return id, err
}

// GetUserById retrieves User by Id and returns error if Id doesn't exist
func GetUserById(id int64) (user *User, err error) {
	ormer := orm.NewOrm()
	user = &User{}

	err = ormer.QueryTable(User{}).Filter("Id", id).RelatedSel().One(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetUserByEmail retrieves User by email and returns error if the email doesn't exist.
func GetUserByEmail(email string) (user *User, err error) {
	ormer := orm.NewOrm()
	user = &User{}

	err = ormer.QueryTable(User{}).Filter("Email", email).RelatedSel().One(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

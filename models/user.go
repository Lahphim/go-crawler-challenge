package models

import (
	"github.com/beego/beego/v2/client/orm"
)

// User : User model
type User struct {
	Base

	Email             string `orm:"unique;size(128)"`
	EncryptedPassword string `orm:"size(128)"`
}

func init() {
	orm.RegisterModel(new(User))
}

// GetUserByEmail retrieves User by email and returns error if the email doesn't exist.
func GetUserByEmail(email string) (user *User, err error) {
	ormer := orm.NewOrm()
	user = &User{Email: email}

	err = ormer.QueryTable(User{}).Filter("Email", email).RelatedSel().One(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// AddUser insert a new User into database and returns last inserted Id on success.
func AddUser(user *User) (id int64, err error) {
	ormer := orm.NewOrm()
	id, err = ormer.Insert(user)

	return id, err
}

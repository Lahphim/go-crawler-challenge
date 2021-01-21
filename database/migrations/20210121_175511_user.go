package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// User_20210121_175511 : DO NOT MODIFY
type User_20210121_175511 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &User_20210121_175511{}
	m.Created = "20210121_175511"

	migration.Register("User_20210121_175511", m)
}

// Up : Run the migrations
func (m *User_20210121_175511) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL(`CREATE TABLE "user"
		(
			id SERIAL,
			email varchar(128) NOT NULL,
			encrypted_password varchar(128) NOT NULL,
			created_at datetime NOT NULL,
			updated_at datetime NOT NULL,
			PRIMARY KEY (id),
			UNIQUE(email)
		);`)
}

// Down : Reverse the migrations
func (m *User_20210121_175511) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL(`DROP TABLE "user"`)
}

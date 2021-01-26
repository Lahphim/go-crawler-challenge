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

	_ = migration.Register("User_20210121_175511", m)
}

// Up : Run the migrations
func (m *User_20210121_175511) Up() {
	m.SQL(`CREATE TABLE "user"
		(
			id SERIAL,
			email varchar(128) NOT NULL,
			encrypted_password varchar(128) NOT NULL,
			created_at timestamp NOT NULL,
			updated_at timestamp NOT NULL,
			PRIMARY KEY (id),
			UNIQUE(email)
		);`)
}

// Down : Reverse the migrations
func (m *User_20210121_175511) Down() {
	m.SQL(`DROP TABLE "user"`)
}

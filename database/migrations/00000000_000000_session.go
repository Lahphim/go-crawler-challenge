package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// Session_00000000_000000 : DO NOT MODIFY
type Session_00000000_000000 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Session_00000000_000000{}
	m.Created = "00000000_000000"

	_ = migration.Register("Session_00000000_000000", m)
}

// Up : Run the migrations
func (m *Session_00000000_000000) Up() {
	m.SQL(`CREATE TABLE "session"
		(
			session_key char(64) NOT NULL,
			session_data bytea,
			session_expiry timestamp NOT NULL,
			CONSTRAINT session_key PRIMARY KEY(session_key)
		);`)
}

// Down : Reverse the migrations
func (m *Session_00000000_000000) Down() {
	m.SQL(`DROP TABLE "session"`)
}

package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// Session_20060102_150405 : DO NOT MODIFY
type Session_20060102_150405 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Session_20060102_150405{}
	m.Created = "20060102_150405"

	_ = migration.Register("Session_20060102_150405", m)
}

// Up : Run the migrations
func (m *Session_20060102_150405) Up() {
	m.SQL(`CREATE TABLE "session"
		(
			session_key char(64) NOT NULL,
			session_data bytea,
			session_expiry timestamp NOT NULL,
			CONSTRAINT session_key PRIMARY KEY(session_key)
		);`)
}

// Down : Reverse the migrations
func (m *Session_20060102_150405) Down() {
	m.SQL(`DROP TABLE "session"`)
}

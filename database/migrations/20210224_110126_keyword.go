package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// Keyword_20210224_110126 : DO NOT MODIFY
type Keyword_20210224_110126 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Keyword_20210224_110126{}
	m.Created = "20210224_110126"

	_ = migration.Register("Keyword_20210224_110126", m)
}

// Up : Run the migrations
func (m *Keyword_20210224_110126) Up() {
	m.SQL(`CREATE TABLE "keyword"
		(
			id SERIAL,
			user_id integer REFERENCES "user" ON DELETE CASCADE,
			keyword varchar(128) NOT NULL,
			url varchar(128) NOT NULL,
			created_at timestamp NOT NULL,
			updated_at timestamp NOT NULL,
			PRIMARY KEY (id)
		);`)
}

// Down : Reverse the migrations
func (m *Keyword_20210224_110126) Down() {
	m.SQL(`DROP TABLE "keyword"`)
}

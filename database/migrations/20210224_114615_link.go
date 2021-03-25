package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// Link_20210224_114615 : DO NOT MODIFY
type Link_20210224_114615 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Link_20210224_114615{}
	m.Created = "20210224_114615"

	_ = migration.Register("Link_20210224_114615", m)
}

// Up : Run the migrations
func (m *Link_20210224_114615) Up() {
	m.SQL(`CREATE TABLE "link"
		(
			id SERIAL,
			keyword_id integer REFERENCES "keyword" ON DELETE CASCADE,
			position_id integer  REFERENCES "position",
			url text NOT NULL,
			created_at timestamp NOT NULL,
			updated_at timestamp NOT NULL,
			PRIMARY KEY (id)
		);`)
}

// Down : Reverse the migrations
func (m *Link_20210224_114615) Down() {
	m.SQL(`DROP TABLE "link"`)
}

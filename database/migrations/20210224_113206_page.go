package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// Page_20210224_113206 : DO NOT MODIFY
type Page_20210224_113206 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Page_20210224_113206{}
	m.Created = "20210224_113206"

	_ = migration.Register("Page_20210224_113206", m)
}

// Up : Run the migrations
func (m *Page_20210224_113206) Up() {
	m.SQL(`CREATE TABLE "page"
		(
			id SERIAL,
			keyword_id integer REFERENCES "keyword" ON DELETE CASCADE,
			raw_html TEXT NOT NULL,
			created_at timestamp NOT NULL,
			updated_at timestamp NOT NULL,
			PRIMARY KEY (id)
		);`)
}

// Down : Reverse the migrations
func (m *Page_20210224_113206) Down() {
	m.SQL(`DROP TABLE "page"`)
}

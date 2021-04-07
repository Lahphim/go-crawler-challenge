package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// Sample_20210112_170701 : DO NOT MODIFY
type Sample_20210112_170701 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Sample_20210112_170701{}
	m.Created = "20210112_170701"

	_ = migration.Register("Sample_20210112_170701", m)
}

// Up : Run the migrations
func (m *Sample_20210112_170701) Up() {
	m.SQL(`CREATE TABLE "sample"
		(
			id SERIAL,
			title varchar(128) NOT NULL,
			description text NOT NULL,
			PRIMARY KEY (id)
		);`)
}

// Down : Reverse the migrations
func (m *Sample_20210112_170701) Down() {
	m.SQL(`DROP TABLE "sample"`)
}

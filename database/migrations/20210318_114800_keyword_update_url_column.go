package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Keyword_20210318_114800 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Keyword_20210318_114800{}
	m.Created = "20210318_114800"

	_ = migration.Register("Keyword_20210318_114800", m)
}

// Up : Run the migrations
func (m *Keyword_20210318_114800) Up() {
	m.SQL(`ALTER TABLE "keyword"
		ALTER COLUMN url DROP NOT NULL
	;`)
}

// Down : Reverse the migrations
func (m *Keyword_20210318_114800) Down() {
	m.SQL(`ALTER TABLE "keyword"
		ALTER COLUMN url SET NOT NULL
	;`)
}

package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Keyword_20210312_173726 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Keyword_20210312_173726{}
	m.Created = "20210312_173726"

	_ = migration.Register("Keyword_20210312_173726", m)
}

// Up : Run the migrations
func (m *Keyword_20210312_173726) Up() {
	m.SQL(`ALTER TABLE "keyword"
		ADD COLUMN status integer DEFAULT 0
	;`)

	m.Update()
}

// Down : Reverse the migrations
func (m *Keyword_20210312_173726) Down() {
	m.SQL(`ALTER TABLE "keyword"
		DROP COLUMN status
	;`)
}

// Update : Update existing records with custom status
func (m *Keyword_20210312_173726) Update() {
	m.SQL(`UPDATE "keyword"
		SET status = 1 where status = 0
	;`)
}

package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// Oauth2_Clients_20210324_125400 : DO NOT MODIFY
type Oauth2_Clients_20210324_125400 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Oauth2_Clients_20210324_125400{}
	m.Created = "20210324_125400"

	_ = migration.Register("Oauth2_Clients_20210324_125400", m)
}

// Up : Run the migrations
func (m *Oauth2_Clients_20210324_125400) Up() {
	m.SQL(`CREATE TABLE "oauth2_clients"
	(
		id TEXT NOT NULL,
		secret TEXT NOT NULL,
		domain TEXT NOT NULL,
		data JSONB NOT NULL,
		PRIMARY KEY (id)
	);`)
}

// Down : Reverse the migrations
func (m *Oauth2_Clients_20210324_125400) Down() {
	m.SQL(`DROP TABLE "oauth2_clients"`)
}

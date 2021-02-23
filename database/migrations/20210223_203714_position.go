package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// Position_20210223_203714 : DO NOT MODIFY
type Position_20210223_203714 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Position_20210223_203714{}
	m.Created = "20210223_203714"

	_ = migration.Register("Position_20210223_203714", m)
}

// Up : Run the migrations
func (m *Position_20210223_203714) Up() {
	m.SQL(`CREATE TABLE "position"
		(
			id SERIAL,
			name varchar(128) NOT NULL,
			selector varchar(128) NOT NULL,
			category varchar(128) NOT NULL,
			created_at timestamp NOT NULL,
			updated_at timestamp NOT NULL,
			PRIMARY KEY (id)
		);`)

	m.Seed()
}

// Down : Reverse the migrations
func (m *Position_20210223_203714) Down() {
	m.SQL(`DROP TABLE "position"`)
}

// Seed : Insert initial data to the database
func (m *Position_20210223_203714) Seed() {
	m.SQL(`INSERT INTO "position"
			(
				name,
				selector,
				category,
				created_at,
				updated_at
			)
			VALUES
				('nonAds', '#search .g .yuRUbf > a', 'normal', current_timestamp, current_timestamp),
				('bottomLinkAds', '#tadsb .d5oMvf > a', 'other', current_timestamp, current_timestamp),
				('otherAds', '#rhs .pla-unit a.pla-unit-title-link', 'other', current_timestamp, current_timestamp),
				('topImageAds', '#tvcap .pla-unit a.pla-unit-title-link', 'top', current_timestamp, current_timestamp),
				('topLinkAds', '#tads .d5oMvf > a', 'top', current_timestamp, current_timestamp);
		`)
}

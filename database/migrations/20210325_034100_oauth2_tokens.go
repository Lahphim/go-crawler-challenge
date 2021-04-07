package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// Oauth2_Tokens_20210325_034100 : DO NOT MODIFY
type Oauth2_Tokens_20210325_034100 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Oauth2_Tokens_20210325_034100{}
	m.Created = "20210325_034100"

	_ = migration.Register("Oauth2_Tokens_20210325_034100", m)
}

// Up : Run the migrations
func (m *Oauth2_Tokens_20210325_034100) Up() {
	m.SQL(`CREATE TABLE "oauth2_tokens"
	(
		id BIGSERIAL NOT NULL,
		created_at TIMESTAMPTZ NOT NULL,
  		expires_at TIMESTAMPTZ NOT NULL,
		code TEXT NOT NULL,
		access TEXT NOT NULL,
		refresh TEXT NOT NULL,
		data JSONB NOT NULL,
		PRIMARY KEY (id)
	);

	CREATE INDEX idx_oauth2_tokens_expires_at ON "oauth2_tokens" (expires_at);
	CREATE INDEX idx_oauth2_tokens_code ON "oauth2_tokens" (code);
	CREATE INDEX idx_oauth2_tokens_access ON "oauth2_tokens" (access);
	CREATE INDEX idx_oauth2_tokens_refresh ON "oauth2_tokens" (refresh);`)
}

// Down : Reverse the migrations
func (m *Oauth2_Tokens_20210325_034100) Down() {
	m.SQL(`DROP TABLE "oauth2_tokens"`)
}

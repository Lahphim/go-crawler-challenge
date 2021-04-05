package oauth

import (
	"github.com/go-oauth2/oauth2/v4/server"
	pg "github.com/vgarvardt/go-oauth2-pg/v4"
)

var (
	ServerOauth *server.Server
	ClientStore *pg.ClientStore
	TokenStore  *pg.TokenStore
)

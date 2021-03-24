package oauth

import (
	"github.com/go-oauth2/oauth2/v4/server"
	pg "github.com/vgarvardt/go-oauth2-pg/v4"
)

type Configuration struct {
	Server      *server.Server
	ClientStore *pg.ClientStore
}

var Server *server.Server
var ClientStore *pg.ClientStore

func (service *Configuration) Run() {
	Server = service.Server
	ClientStore = service.ClientStore
}

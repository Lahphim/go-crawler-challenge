package initializers

import (
	"context"
	"fmt"
	"time"

	"go-crawler-challenge/services/oauth"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/jackc/pgx/v4"
	pg "github.com/vgarvardt/go-oauth2-pg/v4"
	"github.com/vgarvardt/go-pg-adapter/pgx4adapter"
)

func SetUpOauth2() {
	dbURL, err := web.AppConfig.String("dburl")
	if err != nil {
		logs.Critical(fmt.Sprintf("Database URL not found: %v", err))
	}

	pgxConn, err := pgx.Connect(context.TODO(), dbURL)
	if err != nil {
		logs.Critical(fmt.Sprintf("Postgres driver connection failed: %v", err))
	}

	manager := manage.NewDefaultManager()

	// use PostgreSQL token store with pgx.Connection adapter
	adapter := pgx4adapter.NewConn(pgxConn)
	tokenStore, err := pg.NewTokenStore(adapter, pg.WithTokenStoreGCInterval(time.Minute))
	if err != nil {
		logs.Critical(fmt.Sprintf("Create new token store failed: %v", err))
	}
	defer tokenStore.Close()

	clientStore, err := pg.NewClientStore(adapter)
	if err != nil {
		logs.Critical(fmt.Sprintf("Create new client store failed: %v", err))
	}

	// token memory store
	manager.MapTokenStorage(tokenStore)
	// client memory store
	manager.MapClientStorage(clientStore)

	oauthServer := server.NewDefaultServer(manager)
	oauthServer.SetAllowGetAccessRequest(true)
	oauthServer.SetClientInfoHandler(server.ClientFormHandler)

	oauth.ServerOauth = oauthServer
	oauth.ClientStore = clientStore
}

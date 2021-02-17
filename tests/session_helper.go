package tests

import (
	"context"
	"net/http"

	"github.com/beego/beego/v2/server/web"
	"github.com/onsi/ginkgo"
)

// GetSession gets session with given key from cookie, will fail the test if cannot get session store
func GetSession(cookies []*http.Cookie, key string) interface{} {
	c := context.Background()
	for _, cookie := range cookies {
		if cookie.Name == web.BConfig.WebConfig.Session.SessionName {
			store, err := web.GlobalSessions.GetSessionStore(cookie.Value)
			if err != nil {
				ginkgo.Fail("Get store failed" + err.Error())
			}

			return store.Get(c, key)
		}
	}

	return nil
}

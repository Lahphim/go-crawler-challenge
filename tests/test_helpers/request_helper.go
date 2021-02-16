package test_helpers

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"

	"go-crawler-challenge/controllers"
	"go-crawler-challenge/models"

	"github.com/beego/beego/v2/server/web"
	"github.com/onsi/ginkgo"
)

// GetCurrentPath gets current path from HTTP response and return the current url path
func GetCurrentPath(response *http.Response) string {
	url, err := response.Location()
	if err != nil {
		ginkgo.Fail("Get current path failed: " + err.Error())
	}
	return url.Path
}

// MakeRequest makes a HTTP request and returns response
func MakeRequest(method string, url string, body io.Reader) *http.Response {
	request := httpRequest(method, url, body)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	responseRecorder := httptest.NewRecorder()
	web.BeeApp.Handlers.ServeHTTP(responseRecorder, request)

	return responseRecorder.Result()
}

// MakeAuthenticatedRequest makes a HTTP request and returns response by checking with the current session
func MakeAuthenticatedRequest(method string, url string, body io.Reader, user *models.User) *http.Response {
	request := httpRequest(method, url, body)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	responseRecorder := httptest.NewRecorder()
	store, err := web.GlobalSessions.SessionStart(responseRecorder, request)
	if err != nil {
		ginkgo.Fail("Start session failed: " + err.Error())
	}

	err = store.Set(context.Background(), controllers.CurrentUserKey, user.Id)
	if err != nil {
		ginkgo.Fail("Set current user failed: " + err.Error())
	}

	web.BeeApp.Handlers.ServeHTTP(responseRecorder, request)

	return responseRecorder.Result()
}

// httpRequest initiates new HTTP request and handles the error
func httpRequest(method string, url string, body io.Reader) *http.Request {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		ginkgo.Fail("Request failed: " + err.Error())
	}

	return request
}

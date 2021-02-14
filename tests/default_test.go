package test

import (
	"path/filepath"
	"runtime"
	"testing"

	_ "go-crawler-challenge/routers"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	web.TestBeegoInit(apppath)
}

// TestBeego is a sample to run an endpoint test
func TestBeego(t *testing.T) {}

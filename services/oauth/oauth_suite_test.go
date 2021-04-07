package oauth_test

import (
	"testing"

	_ "go-crawler-challenge/conf/initializers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestOauth(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Services/Oauth Suite")
}

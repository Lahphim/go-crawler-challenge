package apiforms_test

import (
	"testing"

	_ "go-crawler-challenge/conf/initializers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestToken(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Forms/Api/Token Suite")
}

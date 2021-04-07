package keyword_test

import (
	"testing"

	_ "go-crawler-challenge/conf/initializers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestKeyword(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Services/Keyword Suite")
}

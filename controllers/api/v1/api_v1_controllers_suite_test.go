package apiv1controllers_test

import (
	"testing"

	_ "go-crawler-challenge/conf/initializers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestV1(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "API V1 Controllers Suite")
}

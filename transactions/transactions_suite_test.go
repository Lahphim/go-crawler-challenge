package transactions_test

import (
	"testing"

	_ "go-crawler-challenge/conf/initializers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTransactions(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Transactions Suite")
}

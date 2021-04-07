package tasks_test

import (
	"context"
	. "go-crawler-challenge/tasks"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SearchKeywordTask", func() {
	Context("#Setup", func() {
		It("runs the task", func() {
			searchKeywordTask := SearchKeywordTask{Name: "search_keyword_task", Schedule: "0 * * * * *"}
			searchKeywordTask.Setup()

			err := searchKeywordTask.Task.Run(context.Background())

			Expect(err).To(BeNil())
		})
	})
})

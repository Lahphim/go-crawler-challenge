package tasks

import (
	"github.com/beego/beego/v2/task"
)

func init() {
	searchKeywordTask := SearchKeywordTask{Name: "search_keyword_task", Schedule: "0 * * * * *"}
	searchKeywordTask.Setup()

	task.AddTask(searchKeywordTask.Name, searchKeywordTask.Task)
	task.StartTask()
}

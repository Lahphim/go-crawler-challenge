package initializers

import (
	. "go-crawler-challenge/tasks"

	"github.com/beego/beego/v2/task"
)

// SetUpTask : Set up all the tasks by scheduler
func SetUpTask() {
	searchKeywordTask := SearchKeywordTask{Name: "search_keyword_task", Schedule: "0 * * * * *"}
	searchKeywordTask.Setup()

	task.AddTask(searchKeywordTask.Name, searchKeywordTask.Task)
	task.StartTask()
}

package tasks

import (
	"context"
	"fmt"
	"go-crawler-challenge/services/scraper"

	"go-crawler-challenge/models"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/task"
)

type SearchKeywordTask struct {
	Name     string
	Schedule string
	Task     *task.Task
}

func (t *SearchKeywordTask) Setup() {
	t.Task = task.NewTask(t.Name, t.Schedule, onScheduledTask)
	logs.Info(fmt.Sprintf("setup task `%v` with schedule at `%v`", t.Name, t.Schedule))
}

func onScheduledTask(_ context.Context) (err error) {
	return nil

	// query an oldest pending status keyword from database
	query := map[string]interface{}{"Status": models.GetStatusKeyword("pending")}
	order := []string{"created_at asc"}
	keyword, err := models.GetKeyword(query, order)
	if err != nil {
		// not found record
		return nil
	}

	searchKeyword := scraper.SearchKeywordService{Keyword: keyword}
	err = searchKeyword.Run()
	if err != nil {
		keyword.Status = models.GetStatusKeyword("failed")

		otherErr := models.UpdateKeyword(keyword)
		if otherErr != nil {
			return otherErr
		}

		return err
	} else {
		keyword.Status = models.GetStatusKeyword("completed")

		err = models.UpdateKeyword(keyword)
		if err != nil {
			return err
		}
	}

	return nil
}

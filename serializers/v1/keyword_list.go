package v1serializers

import (
	"fmt"

	"go-crawler-challenge/helpers"
	"go-crawler-challenge/models"
)

type KeywordList struct {
	KeywordList []*models.Keyword
	TotalRows   int
	PageSize    int
}

type keywordItemResponse struct {
	Id               string `jsonapi:"primary,keywords"`
	Keyword          string `jsonapi:"attr,keyword"`
	Url              string `jsonapi:"attr,url"`
	Status           int    `jsonapi:"attr,status"`
	StatusDetail     string `jsonapi:"attr,status_detail"`
	CreatedAtTimeAgo string `jsonapi:"attr,created_at_time_ago"`
	TotalPages       int    `jsonapi:"meta,total_pages"`
}

func (serializer *KeywordList) Data() (dataList []*keywordItemResponse) {
	for _, keyword := range serializer.KeywordList {
		dataList = append(dataList, &keywordItemResponse{
			Id:               fmt.Sprint(keyword.Id),
			Keyword:          keyword.Keyword,
			Url:              keyword.Url,
			Status:           keyword.Status,
			StatusDetail:     helpers.GetMapKeyByNumber(models.GetKeywordStatuses(), keyword.Status),
			CreatedAtTimeAgo: helpers.ToTimeAgo(keyword.CreatedAt),
		})
	}

	return dataList
}

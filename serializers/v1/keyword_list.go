package v1serializers

import (
	"fmt"

	"go-crawler-challenge/helpers"
	"go-crawler-challenge/models"

	"github.com/beego/beego/v2/adapter/utils/pagination"
	"github.com/google/jsonapi"
)

type KeywordList struct {
	KeywordList []*models.Keyword
	Paginator   *pagination.Paginator
}

type KeywordItemResponse struct {
	Id               string `jsonapi:"primary,keywords"`
	Keyword          string `jsonapi:"attr,keyword"`
	Url              string `jsonapi:"attr,url"`
	Status           int    `jsonapi:"attr,status"`
	StatusDetail     string `jsonapi:"attr,status_detail"`
	CreatedAtTimeAgo string `jsonapi:"attr,created_at_time_ago"`
}

func (serializer *KeywordList) Data() (dataList []*KeywordItemResponse) {
	for _, keyword := range serializer.KeywordList {
		dataList = append(dataList, &KeywordItemResponse{
			Id:               fmt.Sprint(keyword.Id),
			Keyword:          keyword.Keyword,
			Url:              keyword.Url,
			Status:           keyword.Status,
			StatusDetail:     helpers.GetMapKeyByNumberValue(models.GetKeywordStatuses(), keyword.Status),
			CreatedAtTimeAgo: helpers.ToTimeAgo(keyword.CreatedAt),
		})
	}

	return dataList
}

func (serializer *KeywordList) Meta() (meta *jsonapi.Meta) {
	meta = &jsonapi.Meta{
		"total_pages": serializer.Paginator.PageNums(),
	}

	return meta
}

func (serializer *KeywordList) Links() (links *jsonapi.Links) {
	currentPageAt := serializer.Paginator.Page()

	links = &jsonapi.Links{
		"self":  serializer.Paginator.PageLink(currentPageAt),
		"first": serializer.Paginator.PageLinkFirst(),
		"prev":  serializer.Paginator.PageLinkPrev(),
		"next":  serializer.Paginator.PageLinkNext(),
		"last":  serializer.Paginator.PageLinkLast(),
	}

	return links
}

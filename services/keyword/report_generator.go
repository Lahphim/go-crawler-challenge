package keyword

import (
	"go-crawler-challenge/models"

	"github.com/beego/beego/v2/client/orm"
)

type ReportGenerator struct {
	Keyword *models.Keyword
}

// Generate handles processing and generating the report based on a given keyword.
func (service *ReportGenerator) Generate() (reportInterface interface{}, err error) {
	ormer := orm.NewOrm()
	keywordRecord := service.Keyword

	_, err = ormer.LoadRelated(keywordRecord, "Page")
	if err != nil {
		return nil, err
	}

	_, err = ormer.LoadRelated(keywordRecord, "Links")
	if err != nil {
		return nil, err
	}

	linkList, err := initializeLinkList()
	if err != nil {
		return nil, err
	}

	for _, linkRecord := range keywordRecord.Links {
		_, err = ormer.LoadRelated(linkRecord, "Position")
		if err != nil {
			break
		} else {
			linkList[linkRecord.Position.Category] = append(
				linkList[linkRecord.Position.Category],
				linkRecord.Url,
			)
		}
	}

	totalAdsTop := len(linkList["top"])
	totalAdsOther := len(linkList["other"])
	totalNonAds := len(linkList["normal"])

	reportResult := &models.Report{
		Id:      keywordRecord.Id,
		Keyword: keywordRecord.Keyword,
		Url:     keywordRecord.Url,
		RawHtml: keywordRecord.Page.RawHtml,

		CreatedAt: keywordRecord.CreatedAt,
		UpdatedAt: keywordRecord.UpdatedAt,

		LinkList: linkList,

		TotalAdsTop: totalAdsTop,
		TotalAds:    totalAdsTop + totalAdsOther,
		TotalNonAds: totalNonAds,
		TotalLink:   totalAdsTop + totalAdsOther + totalNonAds,
	}

	return reportResult, nil
}

func initializeLinkList() (linkList map[string][]string, err error) {
	linkList = map[string][]string{}
	positions, err := models.GetAllPosition()
	if err != nil {
		return nil, err
	} else {
		for _, positionRecord := range positions {
			linkList[positionRecord.Category] = []string{}
		}
	}

	return linkList, nil
}

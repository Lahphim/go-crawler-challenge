package v1serializers

import (
	"fmt"
	"time"

	"go-crawler-challenge/models"
)

type ReportDetail struct {
	Report *models.Report
}

type reportDetailResponse struct {
	Id      string `jsonapi:"primary,reports"`
	Keyword string `jsonapi:"attr,keyword"`
	Url     string `jsonapi:"attr,url"`
	RawHtml string `jsonapi:"attr,raw_html"`

	LinkList map[string][]string `jsonapi:"attr,link_list"`

	CreatedAt time.Time `jsonapi:"attr,created_at"`
	UpdatedAt time.Time `jsonapi:"attr,updated_at"`

	TotalAdsTop int `jsonapi:"attr,total_ads_top"`
	TotalAds    int `jsonapi:"attr,total_ads"`
	TotalNonAds int `jsonapi:"attr,total_non_ads"`
	TotalLink   int `jsonapi:"attr,total_link"`
}

func (serializer *ReportDetail) Data() (reportDetail *reportDetailResponse) {
	report := serializer.Report

	reportDetail = &reportDetailResponse{
		Id:      fmt.Sprint(report.Id),
		Keyword: report.Keyword,
		Url:     report.Url,
		RawHtml: report.RawHtml,

		LinkList: report.LinkList,

		CreatedAt: report.CreatedAt.Local(),
		UpdatedAt: report.UpdatedAt.Local(),

		TotalAdsTop: report.TotalAdsTop,
		TotalAds:    report.TotalAds,
		TotalNonAds: report.TotalNonAds,
		TotalLink:   report.TotalLink,
	}

	return reportDetail
}

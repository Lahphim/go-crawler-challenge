package v1serializers

type KeywordScraper struct {
	Message string
}

type keywordScraperResponse struct {
	Id      int    `jsonapi:"primary,keyword_scraper"`
	Message string `jsonapi:"attr,message"`
}

func (serializer *KeywordScraper) Data() (data *keywordScraperResponse) {
	data = &keywordScraperResponse{
		Message: serializer.Message,
	}

	return data
}

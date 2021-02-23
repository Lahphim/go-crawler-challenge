package fixtures

import (
	"go-crawler-challenge/models"

	"github.com/onsi/ginkgo"
)

func FabricatePosition(name string, selector string, category string) (position *models.Position) {
	position = &models.Position{
		Name:     name,
		Selector: selector,
		Category: category,
	}
	_, err := models.AddPosition(position)
	if err != nil {
		ginkgo.Fail("Add position failed: " + err.Error())
	}

	return position
}

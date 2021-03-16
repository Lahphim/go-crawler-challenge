package keyword

import (
	"fmt"

	"go-crawler-challenge/models"
	. "go-crawler-challenge/services"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/core/validation"
)

type SearchKeywordService struct {
	Keywords []string
	User     *models.User `valid:"Required;"`
}

func (service *SearchKeywordService) Run() (err error) {
	validator := validation.Validation{}

	valid, err := validator.Valid(service)
	if err != nil {
		return err
	}

	if !valid {
		return validator.Errors[0]
	}

	return err
}

func (service *SearchKeywordService) Valid(validation *validation.Validation) {
	// Validate current existing user
	existingUser, _ := models.GetUserById(service.User.Id)
	if existingUser == nil {
		err := validation.SetError("User", ValidationMessages["InvalidUser"])
		if err == nil {
			logs.Warning(fmt.Sprintf("Set validation error failed: %v", err))
		}
	}

	keywordLength := len(service.Keywords)
	if keywordLength < ContentMinimumSize || keywordLength > ContentMaximumSize {
		err := validation.SetError("Keywords", ValidationMessages["ExceedKeywordSize"])
		if err == nil {
			logs.Warning("Set validation error failed")
		}
	}
}

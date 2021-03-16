package forms

import (
	"mime/multipart"

	. "go-crawler-challenge/forms"
	"go-crawler-challenge/helpers"
	"go-crawler-challenge/models"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/core/validation"
)

type UploadKeywordForm struct {
	File       multipart.File
	FileHeader *multipart.FileHeader
	User       *models.User `valid:"Required"`

	keywordList []string
}

func (form *UploadKeywordForm) Valid(validation *validation.Validation) {
	if form.File == nil {
		err := validation.SetError("File", ValidationMessages["RequireUploadFile"])
		if err == nil {
			logs.Warning("Set validation error failed")
		}

		return
	}

	if !helpers.CheckMatchFileType(form.FileHeader, []string{ContentTypeCSV}) {
		err := validation.SetError("File", ValidationMessages["InvalidUploadFileType"])
		if err == nil {
			logs.Warning("Set validation error failed")
		}
	}

	keywords, err := helpers.ReadCSVFile(form.File)
	if err != nil {
		err := validation.SetError("File", ValidationMessages["OpenUploadFile"])
		if err == nil {
			logs.Warning("Set validation error failed")
		}
	}

	keywordLength := len(keywords)
	if keywordLength <= 0 || keywordLength > 100 {
		err := validation.SetError("File", ValidationMessages["ExceedKeywordSize"])
		if err == nil {
			logs.Warning("Set validation error failed")
		}
	} else {
		form.keywordList = keywords
	}
}

func (form *UploadKeywordForm) Save() (err error) {
	validator := validation.Validation{}

	valid, err := validator.Valid(form)
	if err != nil {
		return err
	}

	if !valid {
		return validator.Errors[0]
	}

	return nil
}

func (form *UploadKeywordForm) GetKeywordList() (keywordList []string) {
	return form.keywordList
}

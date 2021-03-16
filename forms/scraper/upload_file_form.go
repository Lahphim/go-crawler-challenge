package forms

import (
	"mime/multipart"

	. "go-crawler-challenge/forms"
	"go-crawler-challenge/helpers"
	"go-crawler-challenge/models"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/core/validation"
)

type UploadFileForm struct {
	File       multipart.File
	FileHeader *multipart.FileHeader
	User       *models.User `valid:"Required"`

	keywordList []string
}

func (form *UploadFileForm) Valid(validation *validation.Validation) {
	if form.File == nil {
		err := validation.SetError("File", ValidationMessages["RequireFile"])
		if err == nil {
			logs.Warning("Set validation error failed")
		}

		return
	}

	if !helpers.CheckMatchFileType(form.FileHeader, []string{ContentTypeCSV}) {
		err := validation.SetError("File", ValidationMessages["InvalidFileType"])
		if err == nil {
			logs.Warning("Set validation error failed")
		}
	}

	keywordList, err := helpers.ReadFileContent(form.File)
	if err != nil {
		err := validation.SetError("File", ValidationMessages["OpenFile"])
		if err == nil {
			logs.Warning("Set validation error failed")
		}
	}

	contentLength := len(keywordList)
	if contentLength < ContentMinimumSize || contentLength > ContentMaximumSize {
		err := validation.SetError("File", ValidationMessages["ExceedKeywordSize"])
		if err == nil {
			logs.Warning("Set validation error failed")
		}
	} else {
		form.keywordList = keywordList
	}
}

func (form *UploadFileForm) Save() (err error) {
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

func (form *UploadFileForm) GetKeywordList() (keywordList []string) {
	return form.keywordList
}

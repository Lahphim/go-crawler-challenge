package forms

import (
	"mime/multipart"

	"github.com/beego/beego/v2/core/validation"
)

type UploadKeywordForm struct {
	File       multipart.File
	FileHeader *multipart.FileHeader
}

func (form *UploadKeywordForm) Valid() {

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

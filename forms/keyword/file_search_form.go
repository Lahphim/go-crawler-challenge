package forms

import (
	"mime/multipart"

	. "go-crawler-challenge/forms"
	"go-crawler-challenge/helpers"
	"go-crawler-challenge/models"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/core/validation"
)

type FileSearchForm struct {
	File       multipart.File
	FileHeader *multipart.FileHeader
	User       *models.User `valid:"Required"`

	keywordList []string
}

func (form *FileSearchForm) Valid(validation *validation.Validation) {
	if form.File == nil {
		err := validation.SetError("File", ValidationMessages["RequireFile"])
		if err == nil {
			logs.Warning(ValidationMessages["ValidationFailed"])
		}

		return
	}

	if !helpers.CheckMatchFileType(form.FileHeader, []string{KeywordUploadContentTypeCSV}) {
		err := validation.SetError("File", ValidationMessages["InvalidFileType"])
		if err == nil {
			logs.Warning(ValidationMessages["ValidationFailed"])
		}

		return
	}

	keywordList, err := helpers.ReadFileContent(form.File)
	if err != nil {
		err := validation.SetError("File", ValidationMessages["OpenFile"])
		if err == nil {
			logs.Warning(ValidationMessages["ValidationFailed"])
		}

		return
	}

	contentLength := len(keywordList)
	if contentLength < KeywordUploadMinimumSize || contentLength > KeywordUploadMaximumSize {
		err := validation.SetError("File", ValidationMessages["InvalidKeywordSize"])
		if err == nil {
			logs.Warning(ValidationMessages["ValidationFailed"])
		}

		return
	} else {
		form.keywordList = keywordList
	}
}

func (form *FileSearchForm) Save() (err error) {
	validator := validation.Validation{}

	valid, err := validator.Valid(form)
	if err != nil {
		return err
	}

	if !valid {
		return validator.Errors[0]
	}

	err = form.createKeywordList()
	if err != nil {
		return err
	}

	return nil
}

func (form *FileSearchForm) createKeywordList() (err error) {
	ormer := orm.NewOrm()

	// Transaction : Begin
	txnOrmer, err := ormer.Begin()
	if err != nil {
		return err
	}

	// Transaction : Keyword
	var keywordList []models.Keyword
	for _, keywordText := range form.keywordList {
		keyword := models.Keyword{
			Keyword: keywordText,
			User:    form.User,
		}

		keywordList = append(keywordList, keyword)
	}
	_, err = txnOrmer.InsertMulti(50, keywordList)
	if err != nil {
		errRollback := txnOrmer.Rollback()
		if errRollback != nil {
			return errRollback
		}

		return err
	}

	// Transaction : Commit
	err = txnOrmer.Commit()
	if err != nil {
		errRollback := txnOrmer.Rollback()
		if errRollback != nil {
			return errRollback
		}

		return err
	}

	return nil
}

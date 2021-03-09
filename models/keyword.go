package models

import (
	"strings"

	"go-crawler-challenge/helpers"

	"github.com/beego/beego/v2/client/orm"
)

// Keyword : Keyword model
type Keyword struct {
	Base

	User  *User   `orm:"rel(fk)"`
	Page  *Page   `orm:"reverse(one)"`
	Links []*Link `orm:"reverse(many)"`

	Keyword string `orm:"size(128)"`
	Url     string `orm:"size(128)"`
}

func init() {
	orm.RegisterModel(new(Keyword))
}

// GetAllKeyword retrieves all Keyword matches certain condition and returns empty list if no records exist.
func GetAllKeyword(query map[string]interface{}, orderByList []string, offset int64, limit int64) (keywords []*Keyword, err error) {
	ormer := orm.NewOrm()
	querySetter := ormer.QueryTable(Keyword{})

	// query:
	for key, value := range query {
		key = strings.ReplaceAll(key, ".", "__")
		querySetter = querySetter.Filter(key, value)
	}

	// order by:
	querySetter = querySetter.OrderBy(helpers.FormatOrderByFor(orderByList)...).RelatedSel()

	// offset, limit:
	_, err = querySetter.Limit(limit, offset).All(&keywords)
	if err != nil {
		return []*Keyword{}, err
	}

	return keywords, err
}

// CountAllKeyword counts total record for the keyword table
func CountAllKeyword() (totalRows int64, err error) {
	ormer := orm.NewOrm()

	return ormer.QueryTable(Keyword{}).Count()
}

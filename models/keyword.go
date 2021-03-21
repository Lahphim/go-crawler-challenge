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
	Status  int    `orm:"default(0)"`
}

var statusKeyword = map[string]int{"failed": -1, "pending": 0, "completed": 1}

func init() {
	orm.RegisterModel(new(Keyword))
}

// AddKeyword inserts a new Keyword into database and returns last inserted Id on success.
func AddKeyword(keyword *Keyword) (id int64, err error) {
	ormer := orm.NewOrm()
	id, err = ormer.Insert(keyword)

	return id, err
}

// GetKeyword retrieves a Keyword by matching with certain conditions and returning error if it doesn't exist
func GetKeyword(query map[string]interface{}, orderByList []string) (keyword *Keyword, err error) {
	ormer := orm.NewOrm()
	keyword = &Keyword{}
	querySeter := ormer.QueryTable(Keyword{})

	// query:
	for key, value := range query {
		key = strings.ReplaceAll(key, ".", "__")
		querySeter = querySeter.Filter(key, value)
	}

	// order by:
	querySeter = querySeter.OrderBy(helpers.FormatOrderByFor(orderByList)...).RelatedSel()

	err = querySeter.One(keyword)
	if err != nil {
		return nil, err
	}

	return keyword, nil
}

// GetAllKeyword retrieves all Keyword matches certain conditions and returns empty list if no records exist.
func GetAllKeyword(query map[string]interface{}, orderByList []string, offset int64, limit int64) (keywords []*Keyword, err error) {
	ormer := orm.NewOrm()
	querySeter := ormer.QueryTable(Keyword{})

	// query:
	for key, value := range query {
		key = strings.ReplaceAll(key, ".", "__")
		querySeter = querySeter.Filter(key, value)
	}

	// order by:
	querySeter = querySeter.OrderBy(helpers.FormatOrderByFor(orderByList)...).RelatedSel()

	// offset, limit:
	_, err = querySeter.Limit(limit, offset).All(&keywords)
	if err != nil {
		return []*Keyword{}, err
	}

	return keywords, err
}

// CountAllKeyword counts total record for the keyword table
func CountAllKeyword(query map[string]interface{}) (totalRows int64, err error) {
	ormer := orm.NewOrm()
	querySetter := ormer.QueryTable(Keyword{})

	for key, value := range query {
		querySetter = querySetter.Filter(key, value)
	}

	return querySetter.Count()
}

func GetStatusKeyword(status string) int {
	return statusKeyword[status]
}

// UpdateKeyword updates Keyword by Id and returns error if the record to be updated doesn't exist
func UpdateKeyword(keyword *Keyword) (err error) {
	ormer := orm.NewOrm()
	record := Keyword{Base: Base{Id: keyword.Id}}

	err = ormer.Read(&record)
	if err != nil {
		return err
	}

	_, err = ormer.Update(keyword)
	if err != nil {
		return err
	}

	return nil
}

package helpers

import (
	"strings"
)

// BuildOrderByFor builds list of given order list from SQL format to Beego query setter format
// ["user.name", "email desc"] --> ["user__name", "-email"]
func BuildOrderByFor(oldOrderList []string) (newOrderList []string) {
	for _, order := range oldOrderList {
		fieldOrder := strings.Split(order, " ")

		if len(fieldOrder) == 1 {
			newOrderList = append(newOrderList, formattedOrderBy(fieldOrder[0], "asc"))
		} else {
			newOrderList = append(newOrderList, formattedOrderBy(fieldOrder[0], fieldOrder[1]))
		}
	}

	return newOrderList
}

func formattedOrderBy(field string, order string) (newOrderBy string) {
	field = strings.ReplaceAll(field, ".", "__")

	switch strings.ToLower(order) {
	case "desc":
		newOrderBy = "-" + field
	default:
		newOrderBy = field
	}

	return newOrderBy
}

package helpers

import (
	"strings"
)

// FormatOrderByFor builds list of given order list from SQL format to Beego query setter format
// ["user.name", "email desc"] --> ["user__name", "-email"]
func FormatOrderByFor(oldOrderList []string) (formattedOrderParams []string) {
	for _, order := range oldOrderList {
		fieldOrder := strings.Split(order, " ")

		if len(fieldOrder) == 1 {
			formattedOrderParams = append(formattedOrderParams, formattedOrderBy(fieldOrder[0], "asc"))
		} else {
			formattedOrderParams = append(formattedOrderParams, formattedOrderBy(fieldOrder[0], fieldOrder[1]))
		}
	}

	return formattedOrderParams
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

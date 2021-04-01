package helpers

import (
	"math"
)

func GetTotalPages(totalRows int, pageSize int) (totalPages int) {
	return int(math.Ceil(float64(totalRows) / float64(pageSize)))
}

package helpers

// GetMapKeyByNumberValue returns the key of the first entry matched by its value
func GetMapKeyByNumberValue(mapList map[string]int, value int) string {
	for k, v := range mapList {
		if v == value {
			return k
		}
	}

	return ""
}

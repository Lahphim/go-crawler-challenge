package helpers

func GetMapKeyByNumber(mapList map[string]int, value int) string {
	for k, v := range mapList {
		if v == value {
			return k
		}
	}

	return ""
}

package helpers

// GetMapKeyByNumber returns a firstly key from searching by value
func GetMapKeyByNumber(mapList map[string]int, value int) string {
	for k, v := range mapList {
		if v == value {
			return k
		}
	}

	return ""
}

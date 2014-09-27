package chapter1

func HasUniqueChar(value string) bool {
	records := make([]int, 256)

	for _, c := range value {
		index := c - 'a'
		if (records[index] == 0) {
			records[index] = 1
		} else {
			return false
		}
	}	

		return true
}

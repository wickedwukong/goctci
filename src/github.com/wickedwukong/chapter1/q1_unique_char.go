package chapter1

func HasUniqueChar(value string) bool {
	char_set := make([]bool, 256)

	for _, index := range value {
		if (char_set[index]) {
			return false
		} else {
			char_set[index] = true
		}
	}	
		return true
}

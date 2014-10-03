package chapter1

func ReplaceSpaces(value *string) {
	if (len(*value) == 0) {
		return 
	}

	var spacesReplaced []byte

    for i := 0; i < len(*value); i++ {
    	if ((*value)[i] == ' ') {
    		spacesReplaced = append(spacesReplaced, []byte("%20")...)
    	} else {
    		spacesReplaced = append(spacesReplaced, (*value)[i])
    	}
    }

    *value = string(spacesReplaced)
}

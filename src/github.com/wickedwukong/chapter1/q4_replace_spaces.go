package chapter1

func trimSpacesAtBeginning(value *string) {
    index := 0
	for ; index < len(*value) && (*value)[index] == ' '; index++ {
	}

	*value = (*value)[index:]
}

func trimSpacesAtEnd(value *string) {
    index := len(*value) - 1
	for ; index > 0 && (*value)[index] == ' '; index-- {
	}

	*value = (*value)[0:index + 1]
}

func trim(value *string) {
	trimSpacesAtBeginning(value)
	trimSpacesAtEnd(value)
}

func ReplaceSpaces(value *string) {
	if (len(*value) == 0) {
		return 
	}

    trim(value)

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

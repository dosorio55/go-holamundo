package ejercicios

import (
	"strconv"
)

func ReturnValues(text string) (int, string) {
	stringToInteger, error := strconv.Atoi(text)

	if error != nil {
		return 0, "Error, " + error.Error()
	}

	if stringToInteger > 100 {
		return stringToInteger, "value is greater than 100"
	} else if stringToInteger == 100 {
		return stringToInteger, "value is equal to 100"
	} else {
		return stringToInteger, "value is less than 100"
	}

}

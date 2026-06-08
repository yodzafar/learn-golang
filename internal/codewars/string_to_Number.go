package codewars

import "strconv"

func StringToNumber(str string) int {
	result, err := strconv.Atoi(str)

	if err != nil {
		panic(err)
	}

	return result
}

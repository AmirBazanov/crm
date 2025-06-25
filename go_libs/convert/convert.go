package convert

import "strconv"

func StrToInt(str string) (int, error) {
	i, err := strconv.Atoi(str)
	if err != nil {
		return -1, err
	}
	return i, nil
}

func StrToIntWithPanic(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic("String is no numerable" + err.Error())
	}
	return i
}

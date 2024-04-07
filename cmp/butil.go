package cmp

import "strconv"

func i(num int) string {
	return strconv.Itoa(num)
}

func I(num int) string {
	return strconv.Itoa(num)
}

func F(num float64) string {
	return strconv.FormatFloat(num, 'f', -1, 64)
}

package text

import (
	"fmt"
	"strconv"
)

func FloatToStr(somefloat float64) string {
	var out string = fmt.Sprintf("%f", somefloat)
	return out
}
func IntToStr(someInt int) string {
	var out string = strconv.Itoa(someInt)
	return out
}

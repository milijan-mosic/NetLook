package utils

import (
	"fmt"
	"strconv"
)

func FloatToString(num float64) string {
	return fmt.Sprintf("%.4f", num)
}

func IntegerToString(num int) string {
	return strconv.Itoa(num)
}

func Integer64ToString(num int64) string {
	return strconv.FormatInt(num, 10)
}

func BytesToGigabytes() float64 {
	return float64(1024 * 1024 * 1024)
}

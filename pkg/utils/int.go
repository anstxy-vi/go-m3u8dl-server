package utils

import (
	"strconv"
)

func Str2Int(str string, reserve int) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		return reserve
	}
	return i
}

func Str2Float(str string, reserve float64) float64 {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return reserve
	}
	return f
}

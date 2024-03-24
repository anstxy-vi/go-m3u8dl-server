package utils

import (
	"fmt"
	"testing"
)

func TestGetDayRangeTime(t *testing.T) {
	start, end := GetDayRangeTime(-1)
	fmt.Println(start, end)
	t.Log("success")
}

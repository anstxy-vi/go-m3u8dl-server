package utils

import (
	"fmt"
	"testing"
)

func TestRandStringBytes(t *testing.T) {
	t.Log(RandStringBytes(16))
}

func TestRandStringBytesMask(t *testing.T) {
	s := RandString(32)
	fmt.Println(s)
	t.Log(s)
}

func TestRandNumStr(t *testing.T) {
	fmt.Println(RandNumStr(6))
	t.Log(RandNumStr(6))
}

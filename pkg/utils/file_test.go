package utils

import (
	"testing"
)

func TestCWD(t *testing.T) {
	cwd, err := CWD()
	if err != nil {
		t.Error(err)
	}

	if !IsDir(cwd) {
		t.Error("cwd is not dir")
	}

	if !IsExist(cwd) {
		t.Error("cwd is not exist")
	}

	t.Log(cwd)
}

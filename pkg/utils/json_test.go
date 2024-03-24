package utils

import (
	"fmt"
	"testing"
)

type td struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func TestJSON(t *testing.T) {
	tdata := &td{ID: 1, Name: "test"}

	data, err := StringifyJSON(tdata)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(data))

	result := &td{}
	err = ParseJSON(data, result)
	if result == nil || err != nil {
		t.Error(err)
	}
	t.Log(result)
}

func TestStringifyJSON(t *testing.T) {
	data, err := StringifyJSON(nil)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(data))
	t.Log(string(data))
}

func TestParseJSON(t *testing.T) {
	tdata := `"asd"`
	r := ""

	err := ParseJSON([]byte(tdata), &r)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(r)

	t.Log(r)
}

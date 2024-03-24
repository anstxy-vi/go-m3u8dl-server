package utils

import (
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func StringifyJSON(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}

func ParseJSON(data []byte, v interface{}) error {
	return json.Unmarshal(data, &v)
}

func ParseJsonObj(data []byte) (r interface{}, err error) {
	err = json.Unmarshal(data, &r)
	return
}

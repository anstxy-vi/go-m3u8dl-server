package utils

import "strings"

func AllLike(str string) string {
	return "%" + strings.TrimSpace(str) + "%"
}

package utils

import "strings"

func PadLeft(input string, length int, padChar rune) string {
	inputLength := len(input)

	if inputLength >= length {
		// 如果字符串长度大于或等于目标长度，则截断字符串
		return input[:length]
	}

	// 计算需要填充的字符数
	padLength := length - inputLength

	// 使用strings.Repeat创建补齐部分，然后拼接到原始字符串前面
	padString := strings.Repeat(string(padChar), padLength)
	result := padString + input

	return result
}

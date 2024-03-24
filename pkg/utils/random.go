package utils

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const letter2Bytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits

	letter2IdxBits = 6                     // 6 bits to represent a letter index
	letter2IdxMask = 1<<letter2IdxBits - 1 // All 1-bits, as many as letterIdxBits
	letter2IdxMax  = 63 / letter2IdxBits   // # of letter indices fitting in 63 bits
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RangeInt(min, max int) int {
	return rand.Intn(max-min) + min
}

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// see: https://stackoverflow.com/a/31832326
func RandString(n int) string {
	src := rand.NewSource(time.Now().UnixNano())
	sb := strings.Builder{}
	sb.Grow(n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			sb.WriteByte(letterBytes[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return sb.String()
}

func RandString2(n int) string {
	src := rand.NewSource(time.Now().UnixNano())
	sb := strings.Builder{}
	sb.Grow(n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letter2IdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letter2IdxMax
		}
		if idx := int(cache & letter2IdxMask); idx < len(letter2Bytes) {
			sb.WriteByte(letter2Bytes[idx])
			i--
		}
		cache >>= letter2IdxBits
		remain--
	}

	return sb.String()
}

func Rand(n int) int {
	return rand.Intn(n)
}

func RandNumStr(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]string, n)
	for i := range b {
		b[i] = strconv.Itoa(rand.Intn(10))
	}
	return strings.Join(b, "")
}

func CreateRandomPass(size int) string {
	ikind, kinds, result := 3, [][]int{{10, 48}, {26, 97}, {26, 65}}, make([]byte, size)
	for i := 0; i < size; i++ {
		ikind = rand.Intn(3)
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return string(result)
}

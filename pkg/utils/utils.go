package utils

import (
	"compress/gzip"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"regexp"
	"strings"
	"unicode/utf8"
)

func FileMd5(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	h := md5.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}

func BytesMd5(data []byte) string {
	h := md5.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}

// 判断目录是否存在
func IsDirExist(dir string) bool {
	info, err := os.Stat(dir)
	if err != nil {
		return false
	}
	return info.IsDir()
}

// 判断文件是否存在
func IsFileExist(file string) bool {
	info, err := os.Stat(file)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

const COPY_BUFFER_SIZE = 1024 * 1024 * 2

func MoveFile(sourcePath, distPath string) (err error) {
	fi, err := os.Open(sourcePath)
	if err != nil {
		return
	}
	fo, err := os.Create(distPath)
	if err != nil {
		fi.Close()
		return
	}
	buf := make([]byte, COPY_BUFFER_SIZE)
	for {
		count, err := fi.Read(buf)
		if count > 0 {
			cnt, err := fo.Write(buf[:count])
			if err != nil {
				fi.Close()
				fo.Close()
				os.RemoveAll(distPath)
				return err
			}
			if cnt != count {
				fi.Close()
				fo.Close()
				os.RemoveAll(distPath)
				return errors.New("拷贝出错")
			}
		}
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			fi.Close()
			fo.Close()
			os.RemoveAll(distPath)
			return err
		}
	}

	fi.Close()
	fo.Close()
	os.RemoveAll(sourcePath)
	return
}

func GetClientIP(r *http.Request) string {
	ip := ClientIP(r)
	if len(ip) > 0 {
		return ip
	}
	return ClientPublicIP(r)
}

// ClientIP 尽最大努力实现获取客户端 IP 的算法。
// 解析 X-Real-IP 和 X-Forwarded-For 以便于反向代理（nginx 或 haproxy）可以正常工作。
func ClientIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}

	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}

	return ""
}

// ClientPublicIP 尽最大努力实现获取客户端公网 IP 的算法。
// 解析 X-Real-IP 和 X-Forwarded-For 以便于反向代理（nginx 或 haproxy）可以正常工作。
func ClientPublicIP(r *http.Request) string {
	var ip string
	for _, ip = range strings.Split(r.Header.Get("X-Forwarded-For"), ",") {
		ip = strings.TrimSpace(ip)
		if ip != "" {
			return ip
		}
	}

	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		//if !HasLocalIPddr(ip) {
		return ip
		//}
	}

	return ""
}

func Reverse[T any](s []T) []T {
	if len(s) > 1 {
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			s[i], s[j] = s[j], s[i]
		}
	}
	return s
}

// PwdCheck 密码校验
func PwdCheck(password string) (err error) {
	//密码检验

	if len([]rune(password)) < 8 || len([]rune(password)) > 20 {
		return fmt.Errorf("register.password.len.error")
	}
	num := `[0-9]{1}`
	a_z := `[a-z]{1}`
	A_Z := `[A-Z]{1}`
	if match, err := regexp.MatchString(num, password); !match || err != nil {
		return errors.New("register.password.needNum")
	}
	if match, err := regexp.MatchString(a_z, password); !match || err != nil {
		return errors.New("register.password.needLowercase.letters")
	}
	if match, err := regexp.MatchString(A_Z, password); !match || err != nil {
		return errors.New("register.password.needUppercase.letters")
	}
	return
}

func EmailCheck(email string) error {
	if match, err := regexp.MatchString(ALLOW_EMAIL_PATTERN, email); !match || err != nil {
		return errors.New("register.email.notsupport")
	}
	return nil
}

func CompressFile(src, dst string) (err error) {
	f, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open log file: %v", err)
	}
	defer f.Close()

	fi, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("failed to stat log file: %v", err)
	}

	// If this file already exists, we presume it was created by
	// a previous attempt to compress the log file.
	gzf, err := os.OpenFile(dst, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, fi.Mode())
	if err != nil {
		return fmt.Errorf("failed to open compressed log file: %v", err)
	}
	defer gzf.Close()

	gz := gzip.NewWriter(gzf)

	defer func() {
		if err != nil {
			os.Remove(dst)
			err = fmt.Errorf("failed to compress log file: %v", err)
		}
	}()

	if _, err := io.Copy(gz, f); err != nil {
		return err
	}
	if err := gz.Close(); err != nil {
		return err
	}
	if err := gzf.Close(); err != nil {
		return err
	}

	if err := f.Close(); err != nil {
		return err
	}
	return nil
}

// 获取字utf-8前几个字符
func GetRunePreCount(s string, count int) (string, error) {
	l := -1
	for i, r := range s {
		if r == utf8.RuneError {
			return "", errors.New("转换失败")
		}
		l += 1
		if count == l {
			return string(([]byte(s)[:i])), nil
		}
	}
	return s, nil
}

func GetRuneLen(s string) int {
	l := 0
	for _, r := range s {
		if r == utf8.RuneError {
			return -1
		}
		l += 1
	}
	return l
}

func HandleNameMask(str string) string {
	l := GetRuneLen(str)
	if l >= 3 {
		mask := ""
		first := ""
		last := ""
		count := 0
		for i, r := range str {
			if r == utf8.RuneError {
				break
			}
			count += 1
			if count == 2 {
				first = string(([]byte(str))[:i])
			} else if count == l {
				last = string(([]byte(str))[i:])
			} else {
				mask += "*"
			}
		}

		return first + mask + last
	} else if l == 2 {
		_str, _ := GetRunePreCount(str, 1)
		return _str + "***"
	} else {
		return str + "***"
	}
}

func HandleEmailMask(str string) string {
	at := rune('@')
	idx := strings.IndexRune(str, at)
	if idx > 0 {
		sub := string(([]byte(str))[:idx])
		end := len(str)
		idx2 := strings.LastIndex(str, ", +")
		if idx2 > 0 && idx2 > idx {
			end = idx2
		}
		return HandleNameMask(sub) + string(([]byte(str))[idx:end])
	} else {
		return HandleNameMask(str)
	}

}

func GetMapKey[T any, K uint | string | int](m map[K]T) []K {
	keys := []K{}
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func SliceToMap[T any, K uint | string | int](slice []T, getKey func(T) K) map[K]T {
	result := make(map[K]T)

	for _, item := range slice {
		key := getKey(item)
		result[key] = item
	}

	return result
}

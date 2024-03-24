package utils

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
	"crypto/tls"
	"io"
	"net/http"
	"time"

	"github.com/dsnet/compress/brotli"
)

var (
	client = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
			MaxIdleConns:        100,
			MaxConnsPerHost:     100,
			MaxIdleConnsPerHost: 100,
			// Proxy:               proxyFunc,
		},
		Timeout: time.Second * 120,
	}
)

func Request(method, addr, body string, headers map[string]string) (resut string, err error) {
	var buf *bytes.Buffer
	var req *http.Request
	if method == "POST" {
		buf = bytes.NewBufferString(body)
		req, _ = http.NewRequest(method, addr, buf)
	} else {
		req, _ = http.NewRequest(method, addr, nil)
	}

	if err != nil {
		return
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			return
		}
		defer reader.Close()
	case "deflate":
		reader = flate.NewReader(resp.Body)
		if err != nil {
			return
		}
	case "br":
		reader, err = brotli.NewReader(resp.Body, nil)
	default:
		reader = resp.Body
	}
	bs, err := io.ReadAll(reader)
	if err != nil {

		return
	}
	resut = string(bs)
	return
}

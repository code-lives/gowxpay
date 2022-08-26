package common

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

// Setmd5 Md5加密
func Setmd5(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}
func HttpGet(u string) string {
	resp, err := http.Get(u)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(body)
}
func HttpPost(url string, m string, t string, data string) string {
	client := &http.Client{}
	req, err := http.NewRequest(m, url, strings.NewReader(data))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", t)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(body)
}

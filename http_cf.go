package cf

import (
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/jianka/cf/cast"
)

// http请求函数

// 发送GET请求
func Get(url string) (string, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	res, err := client.Get(url)
	if err != nil {
		return "", err
	}
	robots, err2 := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err2 != nil {
		return "", err
	}
	return string(robots), nil
}

// 发送POST请求
// url：         请求地址
// data：        POST请求提交的数据
// contentType： 请求体格式，如：application/json multipart/form-data application/x-www-form-urlencoded text/xml
// content：     请求放回的内容
func Post(urls string, data map[string]interface{}, contentType string) (string, error) {
	var dataStr string
	for k, v := range data {
		dataStr += k + "=" + url.QueryEscape(cast.ToString(v)) + "&"
	}
	dataStr = SubStrLast(dataStr)
	resp, err := http.Post(urls, contentType, strings.NewReader(dataStr))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

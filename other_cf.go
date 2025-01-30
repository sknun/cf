package cf

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
	"unicode"
)

// json返回结果模板
func JsonResult(code int, msg string, d ...map[string]interface{}) map[string]interface{} {
	json := map[string]interface{}{
		"code": code,
		"msg":  msg,
	}
	if len(d) > 0 {
		for _, v := range d {
			for k1, v1 := range v {
				json[k1] = v1
			}
		}
	}
	return json
}

// 取访问路由(去掉多余的参数)
func GetAccessPage(u string) string {
	if u == "" {
		return ""
	}
	arr := strings.Split(u, "?")
	return arr[0]
}

// 取验证路由
func GetAccessVerify(u string, adminPath string) string {
	if u == "" {
		return ""
	}
	arr := strings.Split(u, "?")
	res := strings.Replace(arr[0], adminPath, "", -1)
	res = strings.Replace(res, ".html", "", -1)
	return res
}

// 生成随机数
func MtRand(start int, end int) int {
	rand.Seed(time.Now().UnixMicro())
	return rand.Intn(end-start) + start
}

// 取绝对值
func Abs(n int64) int64 {
	if n < 0 {
		return -n
	}
	return n
}

// base64编码
func SetBase64(data []byte) string {
	encoded := base64.StdEncoding.EncodeToString(data)
	return encoded
}

// base64解码
func GetBase64(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}

// 统计字符数量
// hanCount 汉字数量 lowerCount 小写字母数量 capitalCount 大写字母数量 numberCount数字数量
// tabSpaceCount 空格制表符数量 halfCount 半角字符数量 fullCount 全角字符数量 otherCount 其它字符数量
// s 原字符串 i = 1 中文全角占两个字符
// t 返回统计范围 hanCount, lowerCount, capitalCount, numberCount, tabSpaceCount, halfCount, fullCount, otherCount 默认全取
func CountWords(s string, i int, t ...string) uint32 {
	regLower := "abcdefghijklmnopqrstuvwxyz"
	regCapital := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	regNumber := "0123456789"
	regHalf := `~!@#$%^&*()-=_+[]{}|\;:"',./<>?·`
	regFull := `～！¥………（）——「」【】；‘：“，。《》？`
	regTabSpace := " 	"
	var hanCount, lowerCount, capitalCount, numberCount, tabSpaceCount, halfCount, fullCount, otherCount int
	var count uint32 = 0
	for _, v := range s {
		if unicode.Is(unicode.Scripts["Han"], v) {
			hanCount++
		} else if strings.ContainsRune(regLower, v) {
			lowerCount++
		} else if strings.ContainsRune(regCapital, v) {
			capitalCount++
		} else if strings.ContainsRune(regNumber, v) {
			numberCount++
		} else if strings.ContainsRune(regTabSpace, v) {
			tabSpaceCount++
		} else if strings.ContainsRune(regHalf, v) {
			halfCount++
		} else if strings.ContainsRune(regFull, v) {
			fullCount++
		} else {
			otherCount++
		}
	}
	if len(t) > 0 {
		for _, v := range t {
			switch v {
			case "hanCount":
				count += addcount(hanCount, i)
			case "lowerCount":
				count += uint32(lowerCount)
			case "capitalCount":
				count += uint32(capitalCount)
			case "numberCount":
				count += uint32(numberCount)
			case "tabSpaceCount":
				count += uint32(tabSpaceCount)
			case "halfCount":
				count += uint32(halfCount)
			case "fullCount":
				count += addcount(fullCount, i)
			case "otherCount":
				count += uint32(otherCount)
			}
		}
	} else {
		return uint32(hanCount + lowerCount + capitalCount + numberCount + tabSpaceCount + halfCount + fullCount + otherCount)
	}
	return count
}

func addcount(n, i int) uint32 {
	if i == 1 {
		return uint32(n * 2)
	}
	return uint32(n)
}

// 生成随机字符
func MtRandStr(length int) string {
	arr := []string{"2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f", "g", "h", "k", "m",
		"n", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "A", "B", "C", "D", "E", "F", "G", "H", "K",
		"M", "N", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y"}
	code := ""
	count := len(arr)
	for i := 0; i < length; i++ {
		rand.Seed(time.Now().UnixMicro())
		code += arr[rand.Intn(count-1)]
	}
	return code
}

// 转换驼峰为蛇形
func CamelToSnake(value string) string {
	values := []byte(value)
	var results string

	for i := 0; i < len(values); i++ {
		if 'A' <= values[i] && values[i] <= 'Z' {
			values[i] = values[i] - 'A' + 'a'
			if i != 0 {
				results += "_"
			}
		}
		results += string(values[i])
	}
	return results
}

// 获取本机ip
func GetMyIp() (string, error) {
	apis := []string{
		"https://checkip.amazonaws.com",
		"https://api.ipify.org",
	}

	ipChan := make(chan string, len(apis))
	errChan := make(chan error, len(apis))

	for _, api := range apis {
		go func(url string) {
			ip, err := fetchIP(url)
			if err != nil {
				errChan <- err
				return
			}
			ipChan <- ip
		}(api)
	}

	for i := 0; i < len(apis); i++ {
		select {
		case ip := <-ipChan:
			return ip, nil
		case <-time.After(3 * time.Second):
			return "", fmt.Errorf("请求超时")
		case err := <-errChan:
			fmt.Println("请求失败:", err)
		}
	}
	return "", fmt.Errorf("所有请求失败")
}

func fetchIP(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	ip, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(ip), nil
}
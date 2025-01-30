package cf

// IRIS模板函数

import (
	"errors"
	"html/template"
	"strconv"
	"strings"

	"github.com/jianka/cf/cast"
)

// 字符串转html语言
func Str2html(s string) template.HTML {
	str := template.HTML(s)
	return str
}

// 参数解析 uint64
func GetUint64(name string, value map[string][]string) (uint64, error) {
	if v := value[name]; len(v) > 0 {
		n, err := strconv.ParseUint(v[0], 10, 64)
		return n, err
	}
	return 0, errors.New("参数不存在")
}

// 参数解析 int64
func GetInt64(name string, value map[string][]string) (int64, error) {
	if v := value[name]; len(v) > 0 {
		n, err := strconv.ParseInt(v[0], 10, 64)
		return n, err
	}
	return 0, errors.New("参数不存在")
}

// 参数解析 uint32
func GetUint32(name string, value map[string][]string) (uint32, error) {
	if v := value[name]; len(v) > 0 {
		n, err := strconv.ParseUint(v[0], 10, 32)
		return uint32(n), err
	}
	return 0, errors.New("参数不存在")
}

// 参数解析 int32
func GetInt32(name string, value map[string][]string) (int32, error) {
	if v := value[name]; len(v) > 0 {
		n, err := strconv.ParseInt(v[0], 10, 32)
		return int32(n), err
	}
	return 0, errors.New("参数不存在")
}

// 参数解析 uint
func GetUint(name string, value map[string][]string) (uint, error) {
	if v := value[name]; len(v) > 0 {
		n, err := strconv.ParseUint(v[0], 10, 32)
		return uint(n), err
	}
	return 0, errors.New("参数不存在")
}

// 参数解析 int
func GetInt(name string, value map[string][]string) (int, error) {
	if v := value[name]; len(v) > 0 {
		n, err := strconv.ParseInt(v[0], 10, 32)
		return int(n), err
	}
	return 0, errors.New("参数不存在")
}

// 参数解析 uint8
func GetUint8(name string, value map[string][]string) (uint8, error) {
	if v := value[name]; len(v) > 0 {
		n, err := strconv.ParseUint(v[0], 10, 32)
		return uint8(n), err
	}
	return 0, errors.New("参数不存在")
}

// 参数解析 int8
func GetInt8(name string, value map[string][]string) (int8, error) {
	if v := value[name]; len(v) > 0 {
		n, err := strconv.ParseInt(v[0], 10, 32)
		return int8(n), err
	}
	return 0, errors.New("参数不存在")
}

// 参数解析 字符串
func GetStr(name string, value map[string][]string) string {
	if v := value[name]; len(v) > 0 {
		return v[0]
	}
	return ""
}

// 参数解析 切片 uint64
func GetStrSliceInt(name string, value map[string][]string) []uint64 {
	var slice []uint64
	for k, v := range value {
		if strings.Contains(k, name+"[") {
			slice = append(slice, cast.ToUint64(v[0]))
		}
	}
	return slice
}

// 参数解析 切片 字符串型
func GetStrSliceStr(name string, value map[string][]string) []string {
	var slice []string
	for k, v := range value {
		if strings.Contains(k, name+"[") {
			slice = append(slice, v[0])
		}
	}
	return slice
}

// 宽松的比较
func Compare(a interface{}, b interface{}) bool {
	aInt := cast.ToInt64(a)
	bInt := cast.ToInt64(b)
	return aInt == bInt
}

// 不等于
func Neq(a interface{}, b interface{}) bool {
	aInt := cast.ToInt64(a)
	bInt := cast.ToInt64(b)
	return !(aInt == bInt)
}

// 比较字符串
func EqStr(a interface{}, b interface{}) bool {
	aInt := cast.ToString(a)
	bInt := cast.ToString(b)
	return aInt == bInt
}

// 大于
func Gt(a interface{}, b interface{}) bool {
	aInt := cast.ToInt64(a)
	bInt := cast.ToInt64(b)
	return aInt > bInt
}

// bool值判断
func EqBool(a interface{}) bool {
	return cast.ToBool(a)
}

package cf

// 文本、字符串类函数

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/jianka/cf/cast"
)

// 去除字符串最后一位字符
func SubStrLast(s string) string {
	bt := []rune(s)
	end := len(bt) - 1
	return string(bt[0:end])
}

// 去除字符串第一个字符
func SubStrFirst(s string) string {
	bt := []rune(s)
	end := len(bt)
	return string(bt[1:end])
}

// 如果最后一个字符串是"/"的话，去除字符串最后一位字符，如果不是返回原字符串
func SubStrLast2(s string) string {
	bt := []rune(s)
	if string(bt[len(bt)-1]) == "/" {
		end := len(bt) - 1
		return string(bt[0:end])
	}
	return s
}

// 如果最后一个字符串是"/"的话 返回原字符串，如果不是补位 /
func SubStrComplement(s string) string {
	bt := []rune(s)
	if string(bt[len(bt)-1]) == "/" {
		return s
	}
	return s + "/"
}

// 超出范围的字符串显示...
func SubStrShow(s string, i int) string {
	arr := []rune(s)
	all := 0
	for _, v := range arr {
		if v > 127 {
			all += 2
		} else {
			all++
		}
		if all > i {
			break
		}
	}
	if all > i {
		n := i - 3
		all2 := 0
		var ar []rune
		for _, v := range arr {
			// 截取字符串
			if v > 127 {
				all2 += 2
			} else {
				all2++
			}
			if all2 <= n {
				ar = append(ar, v)
			} else {
				break
			}
		}
		return string(ar) + "..."
	} else {
		return s
	}
}

// 整型64切片转字符串
func SliceUint64ToString(s []uint64) string {
	str := strings.Replace(strings.Trim(fmt.Sprint(s), "[]"), " ", ",", -1)
	return str
}

// 状态转换为显示文本
func StatusForSpanText(i interface{}, s ...[]string) string {
	slice := []string{
		`<span class="layui-badge layui-bg-cyan">Off</span>`,
		`<span class="layui-badge layui-bg-blue">On</span>`,
	}
	ret := ""
	n := cast.ToInt(i)
	if len(s) == 3 && len(s[0]) == len(s[1]) && len(s[1]) == len(s[2]) {
		for k, v := range s[0] {
			if cast.ToInt(v) == n {
				ret = `<span class="layui-badge" style="background-color:` + s[2][k] + `;">` + s[1][k] + `</span>`
				break
			}
		}
	} else {
		if n >= 0 && n < len(slice) {
			ret = slice[n]
		}
	}
	return ret
}

// 状态转换为显示文本 第二版本
func StatusForSpanText2(i interface{}, s ...[]string) string {
	slice := []string{
		`<span class="tag-item-new tag-item-danger">Off</span>`,
		`<span class="tag-item-new">On</span>`,
	}
	ret := ""
	n := cast.ToInt(i)
	if len(s) == 3 && len(s[0]) == len(s[1]) && len(s[1]) == len(s[2]) {
		for k, v := range s[0] {
			if cast.ToInt(v) == n {
				ret = `<span class="tag-item-new" style="background-color:` + s[2][k] + `;">` + s[1][k] + `</span>`
				break
			}
		}
	} else if len(s) == 1 {
		if n >= 0 && n < len(slice) && len(s[0]) == 2 {
			slice2 := []string{
				`<span class="tag-item-new tag-item-danger">` + s[0][0] + `</span>`,
				`<span class="tag-item-new">` + s[0][1] + `</span>`,
			}
			ret = slice2[n]
		}
	} else {
		if n >= 0 && n < len(slice) {
			ret = slice[n]
		}
	}
	return ret
}

// 检测字符串是否为数字
func IsNumeric(s string) bool {
	n := CountWords(s, 1, "numberCount")
	return len(s) == int(n)
}

// 取两个字符串之间值
func ExtractBetween(s, start, end string) string {
	// 找到开始字符串的位置
	startIndex := strings.Index(s, start)
	if startIndex == -1 {
		return ""
	}
	// 找到开始字符串之后的结束字符串的位置
	endStr := s[startIndex+len(start):]
	endIndex := strings.Index(endStr, end)
	if endIndex == -1 {
		return ""
	}
	endIndex += startIndex + len(start)
	// 取得开始和结束字符串之间的内容
	return s[startIndex+len(start) : endIndex]
}

// 判断字符串是否存在 多个值
func CheckStrExists(str string, s ...string) bool {
	if len(s) > 0 {
		for _, v := range s {
			if strings.Contains(str, v) {
				return true
			}
		}
	}
	return false
}

/*
正则批量取文本中间
text 源文本
left 左标识文本
right 右标识文本
isleft 是否不带左边标识 默认带
isright 是否不带右边标识 默认带
*/
func RegexFindAll(text, left, right string, isleft, isright bool) []string {
	// 正则表达式
	regexPattern := left + `([\s\S]*?)` + right
	// 编译正则表达式
	regex, err := regexp.Compile(regexPattern)
	if err != nil {
		return []string{}
	}
	arr := regex.FindAllString(text, -1)
	if isleft {
		for k := range arr {
			arr[k] = strings.Replace(arr[k], left, "", 1)
		}
	}
	if isright {
		for k := range arr {
			arr[k] = strings.Replace(arr[k], right, "", 1)
		}
	}
	// 查找所有匹配项
	return arr
}

// 删除多余字符 空格 空行 制表符
func DeleteExtraCharacters(str string) string {
	// 正则表达式
	regexPattern := `\s+`
	// 编译正则表达式
	re := regexp.MustCompile(regexPattern)
	return re.ReplaceAllString(str, "")
}

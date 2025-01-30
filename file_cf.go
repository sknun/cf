package cf

import (
	"encoding/json"
	"errors"
	"os"
	"time"

	"github.com/jianka/cf/cast"
)

// 读取文件 返回string
func ReadFile(path string) (string, error) {
	f, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(f), nil
}

// 创建文件夹
func CreateFolder(path string) (bool, error) {
	if FileExists(path) {
		return true, nil
	}
	// 不存在时创建文件夹
	if err := os.MkdirAll(path, 0775); err != nil {
		return false, errors.New("创建目录失败")
	}
	return true, nil
}

// 判断所给路径文件/文件夹是否存在
func FileExists(path string) bool {
	//os.Stat获取文件信息
	_, err := os.Stat(path)
	if err != nil {
		if ok := os.IsExist(err); ok {
			return true
		}
		return false
	}
	return true
}

// 取文件大小
func FileSize(path string) int64 {
	//os.Stat获取文件信息
	info, err := os.Stat(path)
	if err != nil {
		return 0
	}
	return info.Size()
}

// 记录日志
// dir 路径
// fileName 文件名
// data 数据
// concise 是否记录时间 false 纯记录数据
// nameRule 名字规则 1 记录全部日志 2 按天分割日志 3 按小时分割日志
// confirm 是否执行记录 默认记录
func RunLog(dir, fileName string, data interface{}, concise bool, nameRule int, confirm ...bool) {
	if len(confirm) == 1 && confirm[0] {
		file := ""
		// 处理文件名
		if nameRule == 1 {
			file = CamelToSnake(fileName) + ".txt"
		} else if nameRule == 2 {
			file = CamelToSnake(fileName) + "-" + time.Now().Format("20060102") + ".txt"
		} else if nameRule == 3 {
			file = CamelToSnake(fileName) + "-" + time.Now().Format("20060102 15") + ".txt"
		}
		// 处理数据
		var txt string
		jsonData, err := json.Marshal(data)
		if err != nil {
			txt = cast.ToString(data)
		} else {
			txt = string(jsonData)
		}
		// 记录时间
		if concise {
			txt = time.Now().String() + "\n" + txt
		}
		f, err := os.OpenFile(SubStrComplement(dir)+file, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0775)
		if err != nil {
			return
		}
		defer f.Close()
		f.WriteString(txt + "\n")
	}
}

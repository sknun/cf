package cf

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

// 计算MD5
func Md5(str string) string {
	m := md5.New()
	m.Write([]byte(str))
	return string(hex.EncodeToString(m.Sum(nil)))
}

// 获取文件md5
func GetFileMD5(pathName string) (string, error) {
	f, err := os.Open(pathName)
	if err != nil {
		return "", err
	}
	defer f.Close()

	md5hash := md5.New()
	if _, err := io.Copy(md5hash, f); err != nil {
		return "", err
	}
	return string(hex.EncodeToString(md5hash.Sum(nil))), nil
}

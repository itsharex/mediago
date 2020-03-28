package utils

import (
	"os"
	"regexp"
)

var (
	urlReg = regexp.MustCompile("^https?://")
)

func CheckDirAndAccess(pathString string) (err error) {
	// 检查下载路径是否存在
	// 并且检查时候有权限写入文件
	fileInfo, err := os.Stat(pathString)
	if err != nil && os.IsNotExist(err) && !fileInfo.IsDir() {
		return
	}
	// fixme： 检查时候有权限写入
	return
}

func IsUrl(str string) bool {
	return urlReg.MatchString(str)
}

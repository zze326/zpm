package util

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"
)

// CheckPathInfo 检查给定路径是否存在，是否是一个目录
func CheckPathInfo(path string) (error, bool, bool) {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, false, false
		}
		// 这里可以根据需要进一步处理其他类型的错误
		fmt.Errorf("检查路径时遇到其他错误: %v\n", err)
	}
	return nil, true, info.IsDir()
}

// PathIsAbsolute 判断给定路径是否为绝对路径
func PathIsAbsolute(pathStr string) bool {
	return path.IsAbs(pathStr) || strings.HasPrefix(pathStr, "/") || strings.HasPrefix(pathStr, "\\") ||
		(len(pathStr) >= 2 && strings.HasPrefix(pathStr, ":\\"))
}

func FatalErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

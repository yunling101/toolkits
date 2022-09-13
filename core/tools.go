package core

import "os"

// 获取环境变量
func GetDefaultEnv(k, v string) string {
	if os.Getenv(k) != "" {
		return os.Getenv(k)
	}
	return v
}

// 创建目录
func Mkdir(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

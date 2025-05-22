package sys

import (
	"bytes"
	"os"
	"os/exec"
)

// GetDefaultEnv 获取环境变量
func GetDefaultEnv(k, v string) string {
	if os.Getenv(k) != "" {
		return os.Getenv(k)
	}
	return v
}

// Mkdir 创建目录
func Mkdir(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

// CmdOut 执行命令
func CmdOut(name string, arg ...string) (string, error) {
	cmd := exec.Command(name, arg...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	return out.String(), err
}

// CmdOutBytes 执行命令
func CmdOutBytes(name string, arg ...string) ([]byte, error) {
	cmd := exec.Command(name, arg...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	return out.Bytes(), err
}

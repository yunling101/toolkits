package file

import (
	"bufio"
	"os"
	"path"
	"sync"
)

var lock = new(sync.RWMutex)

// Basename 获取文件路径中的文件名
func Basename(fp string) string {
	return path.Base(fp)
}

// Create 创建一个新的文件
func Create(name string) (*os.File, error) {
	return os.Create(name)
}

// Remove 删除一个文件
func Remove(name string) error {
	return os.Remove(name)
}

// IsFile 判断是否是一个文件
func IsFile(fp string) bool {
	f, e := os.Stat(fp)
	if e != nil {
		return false
	}
	return !f.IsDir()
}

// IsExist 文件是否存在
func IsExist(fp string) bool {
	_, err := os.Stat(fp)
	return err == nil || os.IsExist(err)
}

// ReadLine 读取一行
func ReadLine(r *bufio.Reader) ([]byte, error) {
	line, isPrefix, err := r.ReadLine()
	for isPrefix && err == nil {
		var bs []byte
		bs, isPrefix, err = r.ReadLine()
		line = append(line, bs...)
	}
	return line, err
}

// ReadAll 文件读取
func ReadAll(filename string) ([]byte, error) {
	lock.Lock()
	defer lock.Unlock()
	return os.ReadFile(filename)
}

// WriteAll 文件写入
func WriteAll(filename string, data []byte) error {
	lock.Lock()
	defer lock.Unlock()
	return os.WriteFile(filename, data, os.ModePerm)
}

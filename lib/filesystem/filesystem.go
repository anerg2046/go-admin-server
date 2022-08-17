package filesystem

import (
	"io/ioutil"
	"os"
)

func CreatePath(path string) (err error) {
	if _, err = os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

// 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return !os.IsNotExist(err)
	}
	return true
}

func CreateFile(path string) error {
	return ioutil.WriteFile(path, []byte{}, 0666)
}

// 删除文件
func Remove(path string) error {
	return os.Remove(path)
}

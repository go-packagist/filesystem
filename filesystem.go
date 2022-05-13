package filesystem

import (
	"io/ioutil"
	"os"
)

func Exists(path string) bool {
	_, err := Stat(path)

	return err == nil
}

func Size(path string) int64 {
	info, err := Stat(path)

	if err != nil {
		return 0
	}

	return info.Size()
}

func IsDir(path string) bool {
	info, err := Stat(path)

	if err != nil {
		return false
	}

	return info.IsDir()
}

func IsFile(path string) bool {
	return !IsDir(path)
}

func FileInfo(path string) (os.FileInfo, error) {
	return Stat(path)
}

func Info(path string) (os.FileInfo, error) {
	return Stat(path)
}

func Stat(path string) (os.FileInfo, error) {
	return os.Stat(path)
}

func ScanDir(path string) ([]os.FileInfo, error) {
	return ReadDir(path)
}

func ReadDir(path string) ([]os.FileInfo, error) {
	return ioutil.ReadDir(path)
}

func ReadFile(path string) (string, error) {
	content, err := ioutil.ReadFile(path)

	if err != nil {
		return "", err
	}

	return string(content), nil
}

func WriteFile(path string, data []byte) error {
	return nil
}

func Mkdir(path string) error {
	return nil
}

func MkdirAll(path string) error {
	return nil
}

func Remove(path string) error {
	return nil
}

func RemoveAll(path string) error {
	return nil
}

func Rename(oldpath, newpath string) error {
	return nil
}

func ReadDirNames(path string) ([]string, error) {
	return nil, nil
}

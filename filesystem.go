package filesystem

import (
	"io/ioutil"
	"os"
)

// Exists checks if a file or directory exists
func Exists(path string) bool {
	_, err := Stat(path)

	return err == nil
}

// Size file size
func Size(path string) int64 {
	info, err := Stat(path)

	if err != nil {
		return 0
	}

	return info.Size()
}

// IsDir checks if a path is a directory
func IsDir(path string) bool {
	info, err := Stat(path)

	if err != nil {
		return false
	}

	return info.IsDir()
}

// IsFile checks if a path is a file
func IsFile(path string) bool {
	info, err := Stat(path)

	if err != nil {
		return false
	}

	return !info.IsDir()
}

// FileInfo file info, alias Stat
func FileInfo(path string) (os.FileInfo, error) {
	return Stat(path)
}

// Info file info, alias Stat
func Info(path string) (os.FileInfo, error) {
	return Stat(path)
}

// Stat file info
func Stat(path string) (os.FileInfo, error) {
	return os.Stat(path)
}

// ScanDir scan directory, alias ReadDir
func ScanDir(path string) ([]os.FileInfo, error) {
	return ReadDir(path)
}

// ReadDir read directory
func ReadDir(path string) ([]os.FileInfo, error) {
	return ioutil.ReadDir(path)
}

// Get get file content, alias ReadFile
func Get(path string) (string, error) {
	return ReadFile(path)
}

// Read read file content, alias ReadFile
func Read(path string) (string, error) {
	return ReadFile(path)
}

// ReadFile read file content
func ReadFile(path string) (string, error) {
	content, err := ioutil.ReadFile(path)

	if err != nil {
		return "", err
	}

	return string(content), nil
}

func WriteFile(path, content string) error {
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

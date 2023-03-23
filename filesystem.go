package filesystem

import (
	"os"
	"path/filepath"
)

// Exists reports whether the named file or directory exists.
func Exists(path string) bool {
	_, err := os.Stat(path)

	return !os.IsNotExist(err)
}

// Missing reports whether the named file or directory does not exist.
func Missing(path string) bool {
	return !Exists(path)
}

// IsDir reports whether the named file is a directory.
func IsDir(path string) bool {
	info, err := os.Stat(path)

	if err != nil {
		return false
	}

	return info.IsDir()
}

// IsFile reports whether the named file is a file.
func IsFile(path string) bool {
	info, err := os.Stat(path)

	if err != nil {
		return false
	}

	return !info.IsDir()
}

// IsAbs reports whether the path is absolute.
func IsAbs(path string) bool {
	return filepath.IsAbs(path)
}

// IsLocal reports whether the path is local.
func IsLocal(path string) bool {
	return filepath.IsLocal(path)
}

func Put(path, contents string, lock bool) error {
	// todo
}

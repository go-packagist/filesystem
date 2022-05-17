package filesystem

import (
	"time"
)

// LocalDrive is a local filesystem drive.
type LocalDrive struct {
	config   *LocalDriveConfig
	prefixer *PathPrefixer
}

// LocalDriveConfig is the configuration for a local filesystem drive.
type LocalDriveConfig struct {
	Root string
}

var _ Drive = (*LocalDrive)(nil)

// NewLocalDrive creates a new local filesystem drive.
func NewLocalDrive(config *LocalDriveConfig) Drive {
	return &LocalDrive{
		config:   config,
		prefixer: NewPathPrefixer(config.Root),
	}
}

// Exists checks if the given path exists.
func (l *LocalDrive) Exists(path string) bool {
	return Exists(l.prefixer.PrefixPath(path))
}

// Get returns the contents of the given path.
func (l *LocalDrive) Get(path string) (string, error) {
	return Get(l.prefixer.PrefixPath(path))
}

// Put writes the given data to the given path.
func (l *LocalDrive) Put(path, contents string) error {
	return Put(l.prefixer.PrefixPath(path), contents)
}

func (l *LocalDrive) Prepend(path, contents string) error {
	return Prepend(l.prefixer.PrefixPath(path), contents)
}

func (l *LocalDrive) Append(path, contents string) error {
	return Append(l.prefixer.PrefixPath(path), contents)
}

func (l *LocalDrive) Delete(path string) error {
	return Delete(l.prefixer.PrefixPath(path))
}

func (l *LocalDrive) Copy(from, to string) error {
	return Copy(l.prefixer.PrefixPath(from), l.prefixer.PrefixPath(to))
}

func (l *LocalDrive) Move(from, to string) error {
	return l.Rename(from, to)
}

func (l *LocalDrive) Rename(from, to string) error {
	return Rename(l.prefixer.PrefixPath(from), l.prefixer.PrefixPath(to))
}

func (l *LocalDrive) Size(path string) (int64, error) {
	return Size(l.prefixer.PrefixPath(path))
}

func (l *LocalDrive) LastModified(path string) (time.Time, error) {
	return LastModified(l.prefixer.PrefixPath(path))
}

func (l *LocalDrive) Files(directory string) (files []string, err error) {
	fileInfos, err := ReadDir(l.prefixer.PrefixPath(directory))

	if err != nil {
		files = []string{}
		return
	}

	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			continue
		}

		files = append(files, directory+"/"+fileInfo.Name())
	}

	return
}

func (l *LocalDrive) AllFiles(directory string) (files []string, err error) {
	fileInfos, err := ReadDir(l.prefixer.PrefixPath(directory))

	if err != nil {
		files = []string{}
		return
	}

	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			dirFiles, dirErr := l.AllFiles(directory + "/" + fileInfo.Name())

			if dirErr != nil {
				return []string{}, dirErr
			}

			files = append(files, dirFiles...)

			continue
		}

		files = append(files, directory+"/"+fileInfo.Name())
	}

	return
}

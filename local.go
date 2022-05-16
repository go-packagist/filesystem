package filesystem

import "time"

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

func NewLocalDrive(config *LocalDriveConfig) Drive {
	return &LocalDrive{
		config:   config,
		prefixer: NewPathPrefixer(config.Root),
	}
}

func (l *LocalDrive) Exists(path string) bool {
	return Exists(l.prefixer.PrefixPath(path))
}

func (l *LocalDrive) Get(path string) (string, error) {
	// TODO implement me
	panic("implement me")
}

func (l LocalDrive) Put(path, contents string) error {
	// TODO implement me
	panic("implement me")
}

func (l LocalDrive) Prepend(path, contents string) error {
	// TODO implement me
	panic("implement me")
}

func (l LocalDrive) Append(path, contents string) error {
	// TODO implement me
	panic("implement me")
}

func (l LocalDrive) Delete(path string) error {
	// TODO implement me
	panic("implement me")
}

func (l LocalDrive) Copy(from, to string) error {
	// TODO implement me
	panic("implement me")
}

func (l LocalDrive) Move(from, to string) error {
	// TODO implement me
	panic("implement me")
}

func (l LocalDrive) Rename(from, to string) error {
	// TODO implement me
	panic("implement me")
}

func (l LocalDrive) Size(path string) (int64, error) {
	return Size(l.prefixer.PrefixPath(path))
}

func (l LocalDrive) LastModified(path string) (time.Time, error) {
	// TODO implement me
	panic("implement me")
}

func (l LocalDrive) Files(directory string) ([]string, error) {
	// TODO implement me
	panic("implement me")
}

func (l LocalDrive) AllFiles(directory string) ([]string, error) {
	// TODO implement me
	panic("implement me")
}

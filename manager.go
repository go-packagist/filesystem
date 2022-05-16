package filesystem

import (
	"fmt"
	"time"
)

type Manager struct {
	config *Config
}

type Config struct {
	Default string
	Disk    map[string]interface{}
}

type Drive interface {
	Exists(path string) bool
	Get(path string) (string, error)
	Put(path, contents string) error
	Prepend(path, contents string) error
	Append(path, contents string) error
	Delete(path string) error
	Copy(from, to string) error
	Move(from, to string) error
	Rename(from, to string) error
	Size(path string) (int64, error)
	LastModified(path string) (time.Time, error)
	Files(directory string) ([]string, error)
	AllFiles(directory string) ([]string, error)
}

var disks map[string]Drive

func NewManager(config *Config) *Manager {
	return &Manager{
		config: config,
	}
}

func (m *Manager) Drive(name string) Drive {
	return m.Disk(name)
}

func (m *Manager) Disk(name string) Drive {
	if disk, ok := disks[name]; ok {
		return disk
	}

	return m.resolve(name)
}

func (m *Manager) resolve(name string) Drive {
	config, err := m.getConfig(name)

	if err != nil {
		panic(err)
	}

	switch config.(type) {
	case *LocalDriveConfig:
		return m.createLocalDrive(config.(*LocalDriveConfig))
	default:
		panic(fmt.Sprintf("Unknown drive type: %s", name))
		return nil
	}
}

func (m *Manager) getConfig(name string) (interface{}, error) {
	if config, ok := m.config.Disk[name]; ok {
		return config, nil
	}

	return nil, fmt.Errorf("config [%s] not found", name)
}

func (m *Manager) createLocalDrive(config *LocalDriveConfig) Drive {
	return NewLocalDrive(config)
}

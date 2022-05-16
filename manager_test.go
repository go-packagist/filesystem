package filesystem

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestManager_Local(t *testing.T) {
	// init
	m := NewManager(&Config{
		Default: "local",
		Disk: map[string]interface{}{
			"local": &LocalDriveConfig{
				Root: "temp",
			},
		},
	})

	// Exists
	assert.True(t, m.Disk("local").Exists("base.txt"))
	assert.False(t, m.Disk("local").Exists("none.txt"))

	// Get

	// Size
	size, _ := m.Disk("local").Size("base.txt")
	assert.Equal(t, int64(4), size)
	size2, size2Err := m.Disk("local").Size("none.txt")
	assert.Equal(t, int64(0), size2)
	assert.Error(t, size2Err)
}

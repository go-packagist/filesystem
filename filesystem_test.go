package filesystem

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setup() {
	file, _ := os.Create(".testdata/test.txt")
	defer file.Close()

	file.Write([]byte("test"))
}

func init() {
	setup()
}

func TestExists(t *testing.T) {
	assert.True(t, Exists(".testdata/test.txt"))
	assert.False(t, Exists(".testdata/test_not_exists.txt"))
}

func TestIsDir(t *testing.T) {
	assert.True(t, IsDir(".testdata"))
	assert.False(t, IsDir(".testdata/test.txt"))
	assert.False(t, IsDir(".testdata/test_not_exists.txt"))
}

func TestIsFile(t *testing.T) {
	assert.True(t, IsFile(".testdata/test.txt"))
	assert.False(t, IsFile(".testdata"))
	assert.False(t, IsFile(".testdata/test_not_exists.txt"))
}

func TestIsAbs(t *testing.T) {
	assert.True(t, IsAbs("/test"))
	assert.False(t, IsAbs("./test"))
	assert.False(t, IsAbs("test"))
}

func TestIsLocal(t *testing.T) {
	assert.True(t, IsLocal("./test"))
	assert.False(t, IsLocal("/test"))
	assert.True(t, IsLocal("test"))
}

package test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssertFile(t *testing.T) {
	// setup
	dire := t.TempDir()
	dest := filepath.Join(dire, "base.extn")
	os.WriteFile(dest, []byte("body"), 0640)

	// success
	AssertFile(t, dest, "body")
}

func TestMockDire(t *testing.T) {
	// success
	dire := MockDire(t)
	assert.NotEmpty(t, dire)

	// confirm - directory contents
	files, _ := mockFS.ReadDir(".")
	for _, file := range files {
		orig := filepath.Join(dire, file.Name())
		bytes, _ := mockFS.ReadFile(file.Name())
		AssertFile(t, orig, string(bytes))
	}
}

func TestMockFile(t *testing.T) {
	// success
	orig := MockFile(t, "base.extn", "body")
	AssertFile(t, orig, "body")
}

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
	dest := filepath.Join(dire, "name.extn")
	os.WriteFile(dest, []byte("body"), 0640)

	// success
	AssertFile(t, dest, "body")
}

func TestMockDire(t *testing.T) {
	// success
	dire := MockDire(t)
	assert.DirExists(t, dire)

	// confirm - directory contents
	for base, body := range MockFiles {
		orig := filepath.Join(dire, base)
		AssertFile(t, orig, body)
	}
}

func TestMockFile(t *testing.T) {
	// success
	orig := MockFile(t, "alpha.extn")
	AssertFile(t, orig, "Alpha note.\n")
}

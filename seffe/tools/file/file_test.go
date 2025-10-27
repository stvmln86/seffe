package file

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/seffe/seffe/tools/test"
)

func TestCreate(t *testing.T) {
	// setup
	dire := t.TempDir()
	dest := filepath.Join(dire, "name.extn")

	// success
	err := Create(dest, "body", 0640)
	test.AssertFile(t, dest, "body")
	assert.NoError(t, err)
}

func TestExists(t *testing.T) {
	// setup
	orig := test.MockFile(t, "alpha.extn")

	// success - true
	okay := Exists(orig)
	assert.True(t, okay)

	// success - false
	okay = Exists("/nope.extn")
	assert.False(t, okay)
}

func TestRead(t *testing.T) {
	// setup
	orig := test.MockFile(t, "alpha.extn")

	// success
	body, err := Read(orig)
	assert.Equal(t, "Alpha note.\n", body)
	assert.NoError(t, err)
}

func TestReextn(t *testing.T) {
	// setup
	orig := test.MockFile(t, "alpha.extn")
	dest := strings.Replace(orig, "alpha.extn", "alpha.test", 1)

	// success
	err := Reextn(orig, ".test")
	assert.NoFileExists(t, orig)
	assert.FileExists(t, dest)
	assert.NoError(t, err)
}

func TestRename(t *testing.T) {
	// setup
	orig := test.MockFile(t, "alpha.extn")
	dest := strings.Replace(orig, "alpha.extn", "test.extn", 1)

	// success
	err := Rename(orig, "test")
	assert.NoFileExists(t, orig)
	assert.FileExists(t, dest)
	assert.NoError(t, err)
}

func TestSearch(t *testing.T) {
	// setup
	orig := test.MockFile(t, "alpha.extn")

	// success - true
	okay, err := Search(orig, "ALPH")
	assert.True(t, okay)
	assert.NoError(t, err)

	// success - false
	okay, err = Search(orig, "nope")
	assert.False(t, okay)
	assert.NoError(t, err)
}

func TestUpdate(t *testing.T) {
	// setup
	orig := test.MockFile(t, "alpha.extn")

	// success
	err := Update(orig, "body", 0640)
	test.AssertFile(t, orig, "body")
	assert.NoError(t, err)
}

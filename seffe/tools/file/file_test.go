package file

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/seffe/seffe/tools/test"
)

func TestList(t *testing.T) {
	// setup
	dire := test.MockDire(t)

	// success
	origs := List(dire, ".extn")
	assert.Equal(t, []string{
		filepath.Join(dire, "alpha.extn"),
		filepath.Join(dire, "bravo.extn"),
	}, origs)
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

	// success
	dest, err := Reextn(orig, ".new-extn")
	assert.NoFileExists(t, orig)
	assert.FileExists(t, dest)
	assert.Contains(t, dest, "alpha.new-extn")
	assert.NoError(t, err)
}

func TestRename(t *testing.T) {
	// setup
	orig := test.MockFile(t, "alpha.extn")

	// success
	dest, err := Rename(orig, "new-name")
	assert.NoFileExists(t, orig)
	assert.FileExists(t, dest)
	assert.Contains(t, dest, "new-name.extn")
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

func TestWrite(t *testing.T) {
	// setup
	orig := test.MockFile(t, "alpha.extn")

	// success
	err := Write(orig, "Body.\n", 0640)
	test.AssertFile(t, orig, "Body.\n")
	assert.NoError(t, err)
}

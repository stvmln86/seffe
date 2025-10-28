package note

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/seffe/seffe/tools/test"
)

func mockNote(t *testing.T) *Note {
	orig := test.MockFile(t, "alpha.extn")
	return New(orig, 0640)
}

func TestNew(t *testing.T) {
	// success
	note := mockNote(t)
	assert.Contains(t, note.Orig, "alpha.extn")
	assert.Equal(t, os.FileMode(0640), note.Mode)
}

func TestDelete(t *testing.T) {
	// setup
	note := mockNote(t)
	orig := note.Orig

	// success
	err := note.Delete()
	assert.NoFileExists(t, orig)
	assert.FileExists(t, note.Orig)
	assert.Contains(t, note.Orig, "alpha.trash")
	assert.NoError(t, err)
}

func TestExists(t *testing.T) {
	// setup
	note := mockNote(t)

	// success
	okay := note.Exists()
	assert.True(t, okay)
}

func TestMatch(t *testing.T) {
	// setup
	note := mockNote(t)

	// success
	okay := note.Match("ALPH")
	assert.True(t, okay)
}

func TestName(t *testing.T) {
	// setup
	note := mockNote(t)

	// success
	name := note.Name()
	assert.Equal(t, "alpha", name)
}

func TestRead(t *testing.T) {
	// setup
	note := mockNote(t)

	// success
	body, err := note.Read()
	assert.Equal(t, "Alpha note.\n", body)
	assert.NoError(t, err)
}

func TestRename(t *testing.T) {
	// setup
	note := mockNote(t)
	orig := note.Orig

	// success
	err := note.Rename("name")
	assert.NoFileExists(t, orig)
	assert.FileExists(t, note.Orig)
	assert.Contains(t, note.Orig, "name.extn")
	assert.NoError(t, err)
}

func TestSearch(t *testing.T) {
	// setup
	note := mockNote(t)

	// success
	okay, err := note.Search("ALPH")
	assert.True(t, okay)
	assert.NoError(t, err)
}

func TestWrite(t *testing.T) {
	// setup
	note := mockNote(t)

	// success
	err := note.Write("Body.\n")
	test.AssertFile(t, note.Orig, "Body.\n")
	assert.NoError(t, err)
}

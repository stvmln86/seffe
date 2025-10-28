package path

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDire(t *testing.T) {
	// success
	dire := Dire("/dire/name.extn")
	assert.Equal(t, "/dire", dire)
}

func TestExtn(t *testing.T) {
	// success - full extension
	extn := Extn("/dire/name.extn")
	assert.Equal(t, ".extn", extn)

	// success - double extension
	extn = Extn("/dire/name.extn.extn")
	assert.Equal(t, ".extn.extn", extn)

	// success - empty extension
	extn = Extn("/dire/name.")
	assert.Equal(t, ".", extn)

	// success - no extension
	extn = Extn("/dire/name")
	assert.Equal(t, "", extn)
}

func TestJoin(t *testing.T) {
	// success
	path := Join("/dire", "name", ".extn")
	assert.Equal(t, "/dire/name.extn", path)
}

func TestMatch(t *testing.T) {
	// success - true
	okay := Match("/dire/name.extn", "NAM")
	assert.True(t, okay)

	// success - false
	okay = Match("/dire/name.extn", "nope")
	assert.False(t, okay)
}

func TestName(t *testing.T) {
	// success - full extension
	name := Name("/dire/name.extn")
	assert.Equal(t, "name", name)

	// success - double extension
	name = Name("/dire/name.extn.extn")
	assert.Equal(t, "name", name)

	// success - empty extension
	name = Name("/dire/name.")
	assert.Equal(t, "name", name)

	// success - no extension
	name = Name("/dire/name")
	assert.Equal(t, "name", name)
}

func TestReextn(t *testing.T) {
	// success
	dest := Reextn("/dire/name.extn", ".new-extn")
	assert.Equal(t, "/dire/name.new-extn", dest)
}

func TestRename(t *testing.T) {
	// success
	dest := Rename("/dire/name.extn", "new-name")
	assert.Equal(t, "/dire/new-name.extn", dest)
}

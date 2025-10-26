package path

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDire(t *testing.T) {
	// success
	dire := Dire("dire/name.extn")
	assert.Equal(t, "dire", dire)
}

func TestExtn(t *testing.T) {
	// success - full extension
	extn := Extn("dire/name.extn")
	assert.Equal(t, ".extn", extn)

	// success - empty extension
	extn = Extn("dire/name.")
	assert.Equal(t, ".", extn)

	// success - no extension
	extn = Extn("dire/name")
	assert.Equal(t, "", extn)
}

func TestJoin(t *testing.T) {
	// success
	path := Join("dire", "name", ".extn")
	assert.Equal(t, "dire/name.extn", path)
}

func TestMatch(t *testing.T) {
	// success - true
	okay := Match("dire/name.extn", "NAME")
	assert.True(t, okay)

	// success - false
	okay = Match("dire/name.extn", "NOPE")
	assert.False(t, okay)
}

func TestName(t *testing.T) {
	// success
	name := Name("dire/name.extn")
	assert.Equal(t, "name", name)
}

func TestRename(t *testing.T) {
	// success
	dest := Rename("dire/name.extn", "test")
	assert.Equal(t, "dire/test.extn", dest)
}

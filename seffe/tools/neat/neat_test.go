package neat

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDire(t *testing.T) {
	// success
	dire := Dire("/././dire")
	assert.Equal(t, "/dire", dire)
}

func TextExtn(t *testing.T) {
	// success
	extn := Extn("\tEXTN\n")
	assert.Equal(t, ".extn", extn)
}

func TestName(t *testing.T) {
	// success
	name := Name("\tNAME_123!!!\n")
	assert.Equal(t, "name-123", name)
}

func TestTime(t *testing.T) {
	// setup
	want, _ := time.Parse("2006-01-02", "2001-02-03")
	want = want.In(time.Local)

	// success
	tobj := Time("2001-02-03")
	assert.Equal(t, want, tobj)
}

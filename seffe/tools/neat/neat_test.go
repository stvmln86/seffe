package neat

import (
	"testing"

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

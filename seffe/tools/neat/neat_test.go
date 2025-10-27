package neat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBody(t *testing.T) {
	// success
	body := Body("\tBody.\n")
	assert.Equal(t, "Body.\n", body)
}

func TestDire(t *testing.T) {
	// success
	dire := Dire("/././dire")
	assert.Equal(t, "/dire", dire)
}

func TestExtn(t *testing.T) {
	// success - full extension
	extn := Extn("\t.EXTN\n")
	assert.Equal(t, ".extn", extn)

	// success - empty extension
	extn = Extn("\t.\n")
	assert.Equal(t, ".", extn)

	// success - no extension
	extn = Extn("\n")
	assert.Equal(t, "", extn)
}

func TestName(t *testing.T) {
	// success
	name := Name("\tNAME_123!!!\n")
	assert.Equal(t, "name-123", name)
}

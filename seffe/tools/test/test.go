// Package test implements unit testing data and functions.
package test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

// MockFiles is a base:body map of mock files for unit testing.
var MockFiles = map[string]string{
	"alpha.extn": "Alpha note.\n",
	"bravo.extn": "Bravo note.\n",
}

// AssertFile asserts a file's body is equal to a string.
func AssertFile(t *testing.T, orig, want string) {
	bytes, err := os.ReadFile(orig)
	assert.Equal(t, want, string(bytes))
	assert.NoError(t, err)
}

// MockDire returns a temporary directory containing all MockFiles entries.
func MockDire(t *testing.T) string {
	dire := t.TempDir()
	for base, body := range MockFiles {
		dest := filepath.Join(dire, base)
		if err := os.WriteFile(dest, []byte(body), 0640); err != nil {
			t.Fatal(err)
		}
	}

	return dire
}

// MockFile returns a temporary file containing a MockFiles entry.
func MockFile(t *testing.T, base string) string {
	dire := t.TempDir()
	dest := filepath.Join(dire, base)
	if err := os.WriteFile(dest, []byte(MockFiles[base]), 0640); err != nil {
		t.Fatal(err)
	}

	return dest
}

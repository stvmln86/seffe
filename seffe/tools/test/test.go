// Package test implements unit testing data and functions.
package test

import (
	"embed"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed *.extn *.json
var mockFS embed.FS

// AssertFile asserts a file's body is equal to a string.
func AssertFile(t *testing.T, orig, want string) {
	bytes, err := os.ReadFile(orig)
	assert.Equal(t, want, string(bytes))
	assert.NoError(t, err)
}

// MockDire returns the path to a temporary directory containing embedded mock files.
func MockDire(t *testing.T) string {
	files, err := mockFS.ReadDir(".")
	if err != nil {
		t.Fatal(err)
	}

	dire := t.TempDir()
	for _, file := range files {
		dest := filepath.Join(dire, file.Name())
		bytes, err := mockFS.ReadFile(file.Name())
		if err != nil {
			t.Fatal(err)
		}

		if err := os.WriteFile(dest, bytes, 0640); err != nil {
			t.Fatal(err)
		}
	}

	return dire
}

// MockFile returns the path to a temporary file containing a string.
func MockFile(t *testing.T, base, body string) string {
	dire := t.TempDir()
	dest := filepath.Join(dire, base)
	if err := os.WriteFile(dest, []byte(body), 0640); err != nil {
		t.Fatal(err)
	}

	return dest
}

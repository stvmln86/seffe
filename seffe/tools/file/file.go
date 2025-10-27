// Package file implements file system I/O functions.
package file

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/stvmln86/seffe/seffe/tools/path"
)

// Create creates a new file with a body string.
func Create(dest, body string, mode os.FileMode) error {
	if err := os.WriteFile(dest, []byte(body), mode); err != nil {
		return fmt.Errorf("cannot create file %q - %w", dest, err)
	}

	return nil
}

// Exists returns true if a file exists.
func Exists(orig string) bool {
	_, err := os.Stat(orig)
	return !errors.Is(err, os.ErrNotExist)
}

// List returns all files in a directory matching an extension.
func List(dire, extn string) []string {
	glob := filepath.Join(dire, "*"+extn)
	origs, _ := filepath.Glob(glob)
	slices.Sort(origs)
	return origs
}

// Read returns an existing file's body as a string.
func Read(orig string) (string, error) {
	bytes, err := os.ReadFile(orig)
	if err != nil {
		return "", fmt.Errorf("cannot read file %q - %w", orig, err)
	}

	return string(bytes), nil
}

// Reextn moves an existing file to a new extension.
func Reextn(orig, name string) error {
	dest := path.Reextn(orig, name)
	if err := os.Rename(orig, dest); err != nil {
		return fmt.Errorf("cannot rename file %q - %w", orig, err)
	}

	return nil
}

// Rename moves an existing file to a new name.
func Rename(orig, name string) error {
	dest := path.Rename(orig, name)
	if err := os.Rename(orig, dest); err != nil {
		return fmt.Errorf("cannot rename file %q - %w", orig, err)
	}

	return nil
}

// Search returns true if a file's body contains a case-insensitive substring.
func Search(orig, text string) (bool, error) {
	bytes, err := os.ReadFile(orig)
	if err != nil {
		return false, fmt.Errorf("cannot search file %q - %w", orig, err)
	}

	text = strings.ToLower(text)
	body := strings.ToLower(string(bytes))
	return strings.Contains(body, text), nil
}

// Update overwrites an existing file's body with a string.
func Update(orig, body string, mode os.FileMode) error {
	if err := os.WriteFile(orig, []byte(body), mode); err != nil {
		return fmt.Errorf("cannot update file %q - %w", orig, err)
	}

	return nil
}

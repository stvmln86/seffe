// Package file implements file system I/O functions.
package file

import (
	"errors"
	"fmt"
	"os"

	"github.com/stvmln86/seffe/seffe/tools/path"
)

// Create creates a new file with a body string.
func Create(dest, body string, mode os.FileMode) error {
	if Exists(dest) {
		return fmt.Errorf("cannot create file %q - already exists", dest)
	}

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

// Read returns an existing file's body as a string.
func Read(orig string) (string, error) {
	if !Exists(orig) {
		return "", fmt.Errorf("cannot read file %q - does not exist", orig)
	}

	bytes, err := os.ReadFile(orig)
	if err != nil {
		return "", fmt.Errorf("cannot read file %q - %w", orig, err)
	}

	return string(bytes), nil
}

// Rename moves an existing file to a new name.
func Rename(orig, name string) error {
	if !Exists(orig) {
		return fmt.Errorf("cannot rename file %q - does not exist", orig)
	}

	dest := path.Rename(orig, name)
	if err := os.Rename(orig, dest); err != nil {
		return fmt.Errorf("cannot rename file %q - %w", orig, err)
	}

	return nil
}

// Update overwrites an existing file's body with a string.
func Update(orig, body string, mode os.FileMode) error {
	if !Exists(orig) {
		return fmt.Errorf("cannot update file %q - does not exist", orig)
	}

	if err := os.WriteFile(orig, []byte(body), mode); err != nil {
		return fmt.Errorf("cannot update file %q - %w", orig, err)
	}

	return nil
}

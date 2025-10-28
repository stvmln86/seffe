// Package note implements the Note type and methods.
package note

import (
	"os"

	"github.com/stvmln86/seffe/seffe/tools/file"
	"github.com/stvmln86/seffe/seffe/tools/neat"
	"github.com/stvmln86/seffe/seffe/tools/path"
)

// Note is a single plaintext note file in a directory.
type Note struct {
	Orig string
	Mode os.FileMode
}

// New returns a new Note.
func New(orig string, mode os.FileMode) *Note {
	return &Note{orig, mode}
}

// Delete moves the Note to a ".trash" extension and updates the Note's path.
func (n *Note) Delete() error {
	dest, err := file.Reextn(n.Orig, ".trash")
	if err != nil {
		return err
	}

	n.Orig = dest
	return nil
}

// Exists returns true if the Note exists.
func (n *Note) Exists() bool {
	return file.Exists(n.Orig)
}

// Match returns true if the Note's name begins with a case-insensitive substring.
func (n *Note) Match(text string) bool {
	return path.Match(n.Orig, text)
}

// Name returns the Note's name.
func (n *Note) Name() string {
	name := path.Name(n.Orig)
	return neat.Name(name)
}

// Read returns the Note's body as a string.
func (n *Note) Read() (string, error) {
	body, err := file.Read(n.Orig)
	return neat.Body(body), err
}

// Rename moves the Note to a new name and updates the Note's path.
func (n *Note) Rename(name string) error {
	name = neat.Name(name)
	dest, err := file.Rename(n.Orig, name)
	if err != nil {
		return err
	}

	n.Orig = dest
	return nil
}

// Search returns true if the Note's body contains a case-insensitive substring.
func (n *Note) Search(text string) (bool, error) {
	return file.Search(n.Orig, text)
}

// Write overwrites the Note's body with a string.
func (n *Note) Write(body string) error {
	body = neat.Body(body)
	return file.Write(n.Orig, body, n.Mode)
}

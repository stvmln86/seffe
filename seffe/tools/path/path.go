// Package path implements file path manipulation functions.
package path

import (
	"path/filepath"
	"strings"
)

// Dire returns a path's parent directory.
func Dire(orig string) string {
	return filepath.Dir(orig)
}

// Extn returns a path's extension with the leading dot.
func Extn(orig string) string {
	base := filepath.Base(orig)
	if clip := strings.Index(base, "."); clip != -1 {
		return base[clip:]
	}

	return ""
}

// Join returns a joined path from a directory, name and extension.
func Join(dire, name, extn string) string {
	return filepath.Join(dire, name+extn)
}

// Match returns true if a path's name contains a case-insensitive substring.
func Match(orig, text string) bool {
	name := Name(orig)
	name = strings.ToLower(name)
	text = strings.ToLower(text)
	return strings.Contains(name, text)
}

// Name returns a path's name without the extension.
func Name(orig string) string {
	base := filepath.Base(orig)
	if clip := strings.Index(base, "."); clip != -1 {
		return base[:clip]
	}

	return base
}

// Rename returns a path with a new name.
func Rename(orig, name string) string {
	dire := Dire(orig)
	extn := Extn(orig)
	return Join(dire, name, extn)
}

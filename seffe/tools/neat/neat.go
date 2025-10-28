// Package neat implements value sanitisation functions.
package neat

import (
	"path/filepath"
	"strings"
	"unicode"
)

// Body returns a whitespace-trimmed body with a trailing newline.
func Body(body string) string {
	return strings.TrimSpace(body) + "\n"
}

// Dire returns a sanitised directory path.
func Dire(dire string) string {
	return filepath.Clean(dire)
}

// Extn returns a lowercase extension with a leading dot, or an empty string.
func Extn(extn string) string {
	extn = strings.ToLower(extn)
	extn = strings.TrimSpace(extn)
	if extn == "" {
		return ""
	}

	return "." + strings.TrimPrefix(extn, ".")
}

// Name returns a lowercase alphanumeric name with dashes.
func Name(name string) string {
	var bldr strings.Builder
	bldr.Grow(len(name))

	for _, char := range strings.ToLower(name) {
		switch {
		case unicode.In(char, unicode.Letter, unicode.Digit):
			bldr.WriteRune(char)
		case unicode.IsSpace(char) || char == '-' || char == '_':
			bldr.WriteRune('-')
		}
	}

	return strings.Trim(bldr.String(), "-")
}

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
	var chars []rune
	for _, char := range strings.ToLower(name) {
		switch {
		case unicode.In(char, unicode.Letter, unicode.Digit):
			chars = append(chars, char)
		case unicode.IsSpace(char) || char == '-' || char == '_':
			chars = append(chars, '-')
		}
	}

	return strings.Trim(string(chars), "-")
}

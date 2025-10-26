// Package neat implements value sanitisation functions.
package neat

import (
	"path/filepath"
	"strings"
	"time"
	"unicode"
)

// Dire returns a sanitised directory path.
func Dire(dire string) string {
	return filepath.Clean(dire)
}

// Extn returns a lowercase extension with the leading dot.
func Extn(extn string) string {
	extn = strings.ToLower(extn)
	extn = strings.TrimSpace(extn)
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

// Time returns a local Time object from a YYYY-MM-DD string.
func Time(date string) time.Time {
	tobj, _ := time.Parse("2006-01-02", date)
	return tobj.In(time.Local)
}

package utils

import "strings"

// RemoveCurlyBraces removes all occurrences of "&{" from the given string.
// If the string does not contain "&{", it returns the original string.
// Each occurrence of "&{" is replaced with the substring from the beginning of the occurrence to the character before the closing curly brace.
// The modified string is returned.
func RemoveCurlyBraces(s string) string {
	if !strings.Contains(s, "&{") {
		return s
	}

	parts := strings.Split(s, "&{")
	for i, part := range parts {
		if len(part) > 0 {
			parts[i] = part[:len(part)-1]
		}
	}

	return strings.Join(parts, "")
}

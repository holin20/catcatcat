package ezgo

import (
	"strings"
	"unicode"
)

func CamelToSnake(camel string) string {
	var result strings.Builder
	for i, r := range camel {
		if unicode.IsUpper(r) {
			// Add an underscore before uppercase letters (if it's not the first letter)
			if i > 0 {
				result.WriteRune('_')
			}
			// Convert the uppercase letter to lowercase
			result.WriteRune(unicode.ToLower(r))
		} else {
			// Append the lowercase letter as-is
			result.WriteRune(r)
		}
	}
	return result.String()
}

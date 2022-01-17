package unittest

import (
	"strings"
)

// FindFirstStringInBracket - func to get string in first bracket
func findFirstStringInBracket(str string) string {
	res := ""
	ok := 0

	// Check string in first bracket
	for _, char := range strings.Split(str, "") {
		if "(" == char {
			ok++
			// Check ok == 1 to skip store first "("
			if ok == 1 {
				continue
			}
		} else if ")" == char && ok > 0{
			// Check ok > 0 to handle ")" before first bracket
			ok--
			// Stop the loop if first bracket closed ")"
			if ok == 0 {
				break
			}
		}
		// Store char in first bracket
		if ok > 0 {
			res += char
		}
	}

	// Check if bracket was closed or string is null
	if ok == 0 {
		return res
	}

	return ""
}
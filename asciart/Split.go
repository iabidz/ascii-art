package asciart

import (
	"strings"
)

// Split_with_new_line splits the input string by newline characters and returns a slice of lines.
func Split_with_new_line(str string) []string {
	word := []string{}
	b := false
	temp_str := ""
	str = strings.ReplaceAll(str, `\n`, "\n") // Replace literal "\n" with actual newline
	for i := 0; i < len(str); i++ {
		if str[i] == '\n' {
			word = append(word, temp_str)
			temp_str = ""
		} else {
			b = true
			temp_str += string(str[i])
		}
	}
	// Add last line if not empty
	if len(temp_str) > 0 {
		word = append(word, temp_str)
		temp_str = ""
	}
	// Handle case where string ends with newline
	if str[len(str)-1] == '\n' && b {
		word = append(word, temp_str)
	}
	return word
}

package asciiart

import (
	"log"
)

func PrintWords(words []string, slice [][]string) string {
	str := ""
	for _, w := range words { // Iterate over each word
		for i := 1; i <= 8; i++ { // Print each line of the ASCII representation
			if len(w) == 0 { // Check if this is a newline ("" means new line)
				str += "\n"
				break // Break to avoid printing multiple newlines
			}
			for _, e := range w {
				if int(e)-32 >= 0 && int(e)-32 <= len(slice)-1 {
					// Print a specific line of each letter's ASCII art
					str += slice[int(e)-32][i]
				} else {
					// If a special character is encountered, return an error
					log.Fatal("Special characters are not allowed.")
				}
			}
			if i < 8 { // After each line, add a newline (except after the last one)
				str += "\n"
			}
		}
	}

	// Remove the last newline if the result does not contain only characters.
	if !ContainChars(str) {
		str = str[:len(str)-1]
	}
	return str
}

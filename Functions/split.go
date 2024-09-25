package asciiart

func Split(str string) []string {
	slice := []string{}
	newStr := ""

	for i := 0; i < len(str); i++ {
		// Check if the current character is '\' followed by 'n'.
		if i < len(str)-1 && str[i] == '\\' && str[i+1] == 'n' {
			// If there's a substring before '\n', add it to the slice.
			if newStr != "" {
				slice = append(slice, newStr)
				newStr = ""
			}
			// Add an empty string in place of '\n' to handle line breaks without duplicating '\n'.
			slice = append(slice, "")
			// Skip the next character ('n').
			i += 1
		} else {
			// Otherwise, append the character to the current substring.
			newStr += string(str[i])
		}
	}

	// If there's any remaining string, add it to the slice.
	if newStr != "" {
		slice = append(slice, newStr)
	}
	return slice
}

package asciiart

func ContainChars(s string) bool {
	// Check if the string contains any printable characters or just newlines.
	for _, r := range s {
		if int(r) >= 32 && int(r) <= 126 {
			return true
		}
	}
	return false
}

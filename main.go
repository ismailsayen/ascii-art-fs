package main

import (
	"fmt"
	"log"
	"os"

	asciiart "asciiart/Functions"
)

func main() {
	banner := "standard"

	// Check if the user has passed exactly one or two arguments.
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
		return
	}
	// Validate the banner style if provided.
	if len(os.Args) == 3 {
		if os.Args[2] != "standard" && os.Args[2] != "shadow" && os.Args[2] != "thinkertoy" {
			fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
			return
		} else {
			banner = os.Args[2]
		}
	}

	// Read symbols from the corresponding file based on the banner style.
	symboles := asciiart.ReadFile(banner)

	// Check that the input string is not empty.
	if len(os.Args[1]) == 0 {
		return
	}

	// Ensure the file contains at least 95 ASCII characters.
	if len(symboles) < 95 {
		log.Fatal("Error: The ASCII art file is incomplete. Ensure all 95 characters are present in the file.")
	}

	// Convert the input string into ASCII art and print the result.
	result := asciiart.PrintWords(asciiart.Split(os.Args[1]), symboles)
	fmt.Println(result)
}

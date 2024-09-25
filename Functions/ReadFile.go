package asciiart

import (
	"bufio"
	"log"
	"os"
)

func ReadFile(name string) [][]string {
	// Opens the file containing the ASCII art symbol definitions.
	file, err := os.Open("./Files/" + name + ".txt")
	if err != nil {
		log.Fatal("Error:", err)
	}

	defer file.Close() // Ensures the file is closed after usage

	scanner := bufio.NewScanner(file) // Initializes a scanner to read the file line by line

	count := 0
	symbol := []string{}
	symbols := [][]string{}

	// Loops through each line of the text file
	for scanner.Scan() {

		symbol = append(symbol, scanner.Text())
		count++

		// Each ASCII symbol in this format consists of 9 lines
		if count == 9 {
			symbols = append(symbols, symbol) // Adds the complete symbol to the list of symbols
			symbol = []string{}
			count = 0
		}
	}
	return symbols
}

package main

import (
	"bufio"
	"fmt"
	"os"

	"asciart/asciart"
)

func main() {
	// Check if the correct number of arguments is provided
	if len(os.Args) != 2 {
		fmt.Println("Usage : go run main.go \"Your text here\"")
		return
	}
	// Try to open the font files in order: standard, shadow, thinkertoy
	file, err := os.Open("standard.txt")
	if err != nil {
		file, err = os.Open("shadow.txt")
		if err != nil {
			file, err = os.Open("thinkertoy.txt")
			if err != nil {
				fmt.Println("Error: failed to open any files")
				return
			}
		}
	}
	// Ensure the file is closed after finishing
	defer file.Close()

	// Create a scanner to read the font file
	Scanner := bufio.NewScanner(file)

	// Parse the ASCII art table from the font file
	asci_table := asciart.ParseAsci(Scanner)

	// Split the input string by new lines
	newstring := asciart.Split_with_new_line(os.Args[1])

	// Print the ASCII art for the input string
	asciart.PrintAsci(newstring, asci_table)
}

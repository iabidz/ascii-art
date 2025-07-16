package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage : go run main.go \"Your text here\"")
		return
	}
	file, err := os.Open("standard.txt")
	if err != nil {
		fmt.Println("error :", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	asci := [][]string{}
	asci_line := []string{}
	
	started := false
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" && !started {
			continue
		}
		if line == "" {
			asci = append(asci, asci_line)
			asci_line = []string{}

		} else {

			asci_line = append(asci_line, line)
			started = true
		}

	}
	fmt.Println("goooooo****")
	fmt.Println(len(asci))
	fmt.Println(len(asci_line))
	fmt.Println(asci_line)
	if len(asci_line) > 0 {
		asci = append(asci, asci_line)
	}

	input := os.Args[1]
	lines := strings.Split(input, "\\n")

	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	for k := 0; k < len(lines); k++ {
		for i := 0; i < 8; i++ {
			for j := 0; j < len(lines[k]); j++ {
				m := (lines[k][j] - 32)
				fmt.Print(asci[m][i])
			}
			fmt.Println()
			if len(lines[k]) == 0 {
				break
			}
		}
	}
}

package asciart

import "fmt"

// PrintAsci prints the ASCII art for each line in str using asci_table.
func PrintAsci(str []string, asci_table [][]string) {
	for k := 0; k < len(str); k++ {
		for i := 0; i < 8; i++ { // Each character is 8 lines tall
			for j := 0; j < len(str[k]); j++ {
				m := (str[k][j] - 32) // Map character to table index
				if m < 94 {
					fmt.Print(asci_table[m][i])
				}

			}
			fmt.Println()
			if len(str[k]) == 0 {
				break // Skip empty lines
			}
		}
	}
}

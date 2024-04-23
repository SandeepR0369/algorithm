package main

import (
	"fmt"
	"strings"
)

func main() {
	// Create a reader to read from standard input
	/*	reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter a word or sentence: ")
		text, _ := reader.ReadString('\n')
	*/
	text := "kytgwu"
	// Normalize the input to lowercase for easier comparison
	text = strings.ToLower(text)

	// Check for the presence of vowels
	foundVowel := false
	for _, char := range text {
		if isVowel(char) {
			foundVowel = true
			break
		}
	}

	// Print the result
	if foundVowel {
		fmt.Println("There are vowels in the entered text.")
	} else {
		fmt.Println("No vowels found in the entered text.")
	}
}

// isVowel checks if a character is a vowel
func isVowel(c rune) bool {
	//fmt.Println("C", string(c))
	switch c {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	default:
		return false
	}
}

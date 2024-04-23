package main

import (
	"fmt"
	"strings"
)

func main() {
	// Create a reader to read from standard input
	/*reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a word or sentence: ")
	text, _ := reader.ReadString('\n')
	*/

	text := "kytgw is a good boy"
	// Normalize the input to lowercase for easier comparison
	text = strings.ToLower(text)

	// Define vowels
	vowels := "aeiou"

	// Check for the presence of vowels
	foundVowel := false
	for _, char := range text {
		//fmt.Println("AA", char, reflect.TypeOf(char), text, "BB", string(char))
		if strings.ContainsRune(vowels, char) {
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

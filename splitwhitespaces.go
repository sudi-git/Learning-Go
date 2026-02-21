package main
//import "fmt"
//func main() {
//fmt.Println(SplitWhitespaces("Hello! How have you been?"))
//}

func SplitWhitespaces(s string) []string {

	// Slice to store all the words
	var final []string

	// Temporary slice to build one word character by character
	var temp []rune

	// Go through the string one character at a time
	for i := 0; i < len(s); i++ {

		// If the character is a space, tab, or newline
		if s[i] == ' ' || s[i] == '\t' || s[i] == '\n' {

			// If a word was built, save it
			if len(temp) > 0 {
				final = append(final, string(temp))

				// Clear temp for the next word
				temp = nil
			}

		} else {
			// Add the character to the current word
			temp = append(temp, rune(s[i]))
		}
	}

	// Add the last word if it exists
	if len(temp) > 0 {
		final = append(final, string(temp))
	}

	// Return the slice of words
	return final
}

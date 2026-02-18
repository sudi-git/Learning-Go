package main
//import "fmt"
//func main() {
	//fmt.Println(ToCapitalCase("Hello! Great to see you! How-are-you-doing-2day?"))
//}

func ToCapitalCase(s string) string {

	// Convert the string into a slice of runes.
		stringCapital := []rune(s)

		
		startLetter := true

		// Go through each character in the string
		for i := 0; i < len(s); i++ {

			// to check if the current character is: (a–z),(A–Z),(0–9)
			if (stringCapital[i] >= 'a' && stringCapital[i] <= 'z') ||(stringCapital[i] >= 'A' && stringCapital[i] <= 'Z') ||
				(stringCapital[i] >= '0' && stringCapital[i] <= '9') {

				// 
				if startLetter {

					//to convert lowercase letter to uppercase
					if stringCapital[i] >= 'a' && stringCapital[i] <= 'z' {
						stringCapital[i] -= 32 // Example: 'a' becomes 'A'
					}

					// 
					startLetter = false

				} else {
					// If the character is NOT the first letter of a word, we make it lowercase
					if stringCapital[i] >= 'A' && stringCapital[i] <= 'Z' {
						stringCapital[i] += 32 // Example: 'A' becomes 'a'
					}
				}

			} else {
				// If the character is NOT a letter or number (space, symbol, punctuation, etc.) the next valid character will start a new word
				startLetter = true
			}
		}

		// it will converts the rune slice back into a string and return it
		return string(stringCapital)
}
	

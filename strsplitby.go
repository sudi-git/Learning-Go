package main
//import "fmt"
//func main() {
//	s := "HowYOUhaveYOUyouYOUbeen?"
//	result := StrSplitBy(s, "YOU")
//	fmt.Println(result)
//}


func StrSplitBy(s, sep string) []string {


		// If the string or separator is empty, return nil
		if sep == "" || s == "" {
			return nil
		}

		// newS will store the final split strings
		var final []string

		// temp will build the current substring character by character
		var temp string

		// Loop through each character in the string
		for i := 0; i < len(s); i++ {

			// Check if the substring starting at i matches the separator
			if i < len(s)-len(sep) && s[i:i+len(sep)] == sep {

				// If separator is found, save the current temp string
				final = append(final, temp)

				// Reset temp for the next substring
				temp = ""

				// Skip ahead by the length of the separator
				i += len(sep)
			}

			// Add the current character to temp
			temp += string(s[i])
		}

		// Append the last collected substring
		final = append(final, temp)

		// Return the result
		return final
}

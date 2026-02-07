package sprint 
func AlphabetMastery(n int) string {
letter:= ""
	for  i:= 0; i <n ;i++{
//changing rune values to string and storing the values in letter
		letter = letter+ string('a'+ rune(i))
	} 
	return letter
}    

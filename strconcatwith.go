package main
//import "fmt"
//func main(){
//toConcat := []string{"Three", " Two", " One", " Go!"}
//fmt.Println(StrConcatWith(toConcat, "."))
//}

func StrConcatWith(strs []string, sep string) string {

	var result string //to strore the combined string

	for _, r := range strs { //to check each string in the slice
		result += r + sep //adding string and septerator
        
	}
	//to check if result if there is any left
		if len(result) > 0 {
		result = result[:len(result)-len(sep)] //removing the last sep from the result 
		}
	
	return result
}

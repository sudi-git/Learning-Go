//package main
//import "fmt"
package sprint
func ReverseAlphabet(step int) string {
	if step <=0 {
		step = 1
	}
		result:= ""
		//start from z ,move backward to a
	for i:= 'z'; i >= 'a'; i -= rune(step){
		//changing rune to string
		result = result + string(rune(i))
	}

	return result
}
//func main() {
	//fmt.Println(ReverseAlphabet(5))
//}

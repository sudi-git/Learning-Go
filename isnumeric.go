package main
//import "fmt"
//func main() {
//	fmt.Println(IsNumeric("01+23-45"))
//}

func IsNumeric(s string) bool {
	//checks every elements of string
	for _, ele:= range s {
		//if the elemets are not a-z
		if !(ele>='0' && ele<='9') {
		return false
		}
	}
return true

}

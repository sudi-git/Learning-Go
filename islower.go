package main
//import "fmt"
//func main() {
//	fmt.Println(IsLower("World!"))
//}

func IsLower(s string) bool {
	//checks every elements of string
	for _, ele:= range s {
		//if the elemets are not a-z
		if !(ele>='a' && ele<='z') {
		return false 
		}
	}
return true

}


package main
//import "fmt"
//func main() {
//fmt.Println(ToUpperCase("Hello! How's your day going?"))
//}

func ToUpperCase(s string) string {
upper:=""
diff:='a'-'A'// to find the difrrence from ASCII a-A
//check every character in the string
	for _, ele:=range s {
		//check if the character is lowercase
		if ele >= 'a' && ele <= 'z'{
			//converts to uppercase
			upper = upper+string(ele-diff)
		} else{
			//to keep character as it is if its not lowercase
			upper = upper+string(ele)
		}
	}
	return upper
}

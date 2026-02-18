package main
//import "fmt"
//func main() {
	//fmt.Println(StrReverse("hello world "))
//}

func StrReverse(s string) string {
result:= ""
//to get the the  index of thr last character
indexn:= (len(s)-1)
	for i:= indexn; i>=0; i--{
		//adding each character to the string
		result = result + string(s[i])
	}
	return result
}

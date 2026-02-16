package main
//import "fmt"
//func main() {
//	fmt.Println(StrLength("Hello World! "))
//}

func StrLength(s string) []int {
//convert string to rune and count the characters
	norune:= []rune(s)
//to know the number of bytes in string
	byt:= len (s)

result:= []int{len(norune), byt}
return result
}

package main
//import "fmt"
//func main() {
//fmt.Println(StrCompare("Hi!", "Hi!"))
//fmt.Println(StrCompare("Day", "ay"))
//fmt.Println(StrCompare("weekday", "week"))
//}

func StrCompare(a, b string) int {
	if a<b {
		return -1
	}else if a>b{
		return 1
	}else{
		return 0
	}
	
}
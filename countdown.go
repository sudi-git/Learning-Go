package main
import "fmt"

func main() {
	fmt.Println(Countdown(7))
	}

//package sprint
func Countdown(n int) string {
	result:= ""

	for i:=n; i>0; i-= 2 {
		//converts int to a rune and then to string
			result = result + string(rune(i) + '0')
			// addd , and space
			 result = result +", "
}
//to add 0! after the result is ready
	return result +"0!"

}

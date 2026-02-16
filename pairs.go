package main
import "fmt"

func main() {
	fmt.Println(Pairs())
}
//package sprint
// 
//import "fmt"
func Pairs() string {
	digit := ""
	for i := 0; i< 99; i++ {
		for j:= i+1; j<100; j++ {
			if i == 98{
			digit = digit +fmt.Sprintf("%02d %02d", i, j)
				return digit
			}else{
			fmt.Sprintf("%02d %02d, ", i, j)
			digit = digit +fmt.Sprintf("%02d %02d, ", i, j)
			}
		}
	}
	return digit
}

package main
import "fmt"

func main() {
	fmt.Println(Combinations())
}
package sprint
import "fmt"
func Combinations() string {
	digits:= ""
	result := ""
	for i:=0; i<8; i++ {
		for j:=i+1; j<9 ;j++ {
			for k:=j+1; k<10; k++ {
						digits = fmt.Sprintf("%d%d%d", i,j,k)

							if result != "" {
								result = result +", "
							}
							result = result + digits
		}
	}
}
return result
}

//package main
//import "fmt"

//func main() {
	//fmt.Println(GenerateRange(-1, 4))
//}

package sprint
func GenerateRange(min, max int) []int {
	if  min>=max {
		return nil
	}
	// makes a slice length 0 to start it empty instead of containing 0
		output:=make([]int, 0)
	for i:=min; i<max ;i++{
		// adds the number i to the slice
		output = append(output, i)
	}
return output

}

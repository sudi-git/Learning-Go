//package main
//import "fmt"
//func main() {
//	fmt.Println(RemoveElementsInRange([]float64{10, .8, -.4, 20, 7.7, 3}, 4, 1))
//}

package sprint
func RemoveElementsInRange(arr []float64, from, to int) []float64 {

	//swap from as to and to as from
	if from > to{
		from ,to = to ,from
	}
	//check if from is negative. if it is from should be 0
	if from <0{
		from=0
	}
	//if to is larger than the capacity of the array
	i:= len(arr)
	if to > i{
		to=i
	}
	// slicng the array
	a:=arr[:from]
	b:=arr[to:]
	//combining slices a and b into a single slice
	result := append(a, b...)
	return result
}

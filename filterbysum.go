package main
//import "fmt"
//func main(){
//fmt.Println(FilterBySum([][]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}}, 9)) }

func FilterBySum(arr [][]int, limit int) [][]int {
var result [][] int //to keep the value passing through loop

// to check the subarray
	for _, row:= range arr {
		var sum int
		// to take the value of each rows elements and find the sum
		// fmt.Println("row, ", row)
		for _, value:= range row {
			// fmt.Println("value ", value)
			sum = sum + value
		}
		if sum >=limit {
			//takes the row(subarray)
			result = append(result,row)
		}
		// fmt.Println("sum, ", sum)
	}
	return result
}

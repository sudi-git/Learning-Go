package main

//import "fmt"

//func main() {
//	fmt.Println(SortIntegerTable([]int{2, 0, 5, 4, 1, 3}))
//}

func SortIntegerTable(table []int) []int {

	// making a new slice with the same length as table so that we can use this to check the elements
	arr := make([]int, len(table))
	copy(arr, table)

	// to keep the checked elements in order
	var result []int
	// Looping until all elements are removed from array
	for len(arr) > 0 {

		// to start from the smallest number
		min := arr[0]
		//to keep the smallest number in index 0
		index := 0

		// to check every elements of the array
		for i, v := range arr {
			//to find the smallest number
			if v < min {
				min = v   // to update smallest value
				index = i // to save its position
			}
		}

		// to keep the smallest value to the result slice
		result = append(result, min)

		// removes the number from the array which has passed and kept in the new array.
		arr = append(arr[:index], arr[index+1:]...)
	}

	return result
}





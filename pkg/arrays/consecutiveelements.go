package main

import "fmt"

/**
 * http://www.geeksforgeeks.org/check-if-array-elements-are-consecutive/
 */
func areElementsConsecutive(arr []int) bool {

	min := getMinElement(arr)
	max := getMaxElement(arr)

	if max-min+1 == len(arr) {

		visited := make([]bool, len(arr))
		for i := 0; i < len(arr); i++ {
			if visited[arr[i]-min] != false {
				return false
			}
			visited[arr[i]-min] = true
		}
		return true
	}
	return false
}

func getMinElement(arr []int) int {

	min := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min
}

func getMaxElement(arr []int) int {

	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

func main() {
	arr := []int{2, 3, 2, 5}
	areElementsConsecutive := areElementsConsecutive(arr)
	if areElementsConsecutive {
		fmt.Println(" array elements are consecutive")
	} else {
		fmt.Println(" array elements are non-consecutive")
	}
}

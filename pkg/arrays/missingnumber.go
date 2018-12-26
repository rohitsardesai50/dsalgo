package main

import "fmt"

/**
https://practice.geeksforgeeks.org/problems/missing-number-in-array/0

Missing number in array
 */
func missingNumInArray(a []int, n int) int {

	x1 := a[0]
	x2 := 1


	for i := 1; i < n; i++ {
		x1 = x1 ^ a[i]
	}

	for i := 2; i <= n+1; i++ {
		x2 = x2 ^ i
	}

	return x1 ^ x2

}

func main() {
	arr := []int{1,2,3,4,6}

	maxSum := missingNumInArray(arr, 6)
	fmt.Printf("Max sum : %d\n",maxSum)
}

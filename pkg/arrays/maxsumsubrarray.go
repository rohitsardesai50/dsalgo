package main

import "fmt"

/**
https://practice.geeksforgeeks.org/problems/kadanes-algorithm/0

Kadane's Algorithm
 */
func maxSumSubArray(a []int) int {

	currMax := a[0]
	maxSoFar := a[0]

	for i := 1; i < len(a); i++ {
		currMax = max(a[i], currMax + a[i])
		maxSoFar = max(maxSoFar, currMax)
		fmt.Printf("curSum : %d, maxSoFar : %d\n", currMax, maxSoFar)
	}

	return maxSoFar
}

func max(a ,b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}


func main() {
	arr := []int{-2, -3, 4, -1, -2, 1 , 5, 3}

	maxSum := maxSumSubArray(arr)
	fmt.Printf("Max sum : %d\n",maxSum)
}

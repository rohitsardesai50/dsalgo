package main

import "fmt"

func findNumberNeededForAnagram(a string, b string) int {

	for i, c := range a {
		fmt.Printf("i%d r %c\n", i, c)
	}
}

func main() {
	arr := []int{1,2,3,4,6}

	maxSum := missingNumInArray(arr, 6)
	fmt.Printf("Max sum : %d\n",maxSum)
}




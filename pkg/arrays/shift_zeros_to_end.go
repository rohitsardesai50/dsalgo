package main

import "fmt"

/**
  https://www.geeksforgeeks.org/move-zeroes-end-array/

  Move all zeroes to end of array
 */
func moveAllZerosToEnd(a []int) []int {

	b := a[:0]
	count := 0
	for _, x := range a {
		if x != 0 {
			b = append(b, x)
			count ++
		}
	}
	for count < len(a) {
		b = append(b , 0)
		count++
	}
	return b
}

func main() {
	arr := []int{1,2,0,9,0,6,0,3,4,6}
	res := moveAllZerosToEnd(arr)
	fmt.Printf("Result : %v\n",res)
}

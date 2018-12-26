package main

import (
	"dsalgo/pkg/stack"
	"fmt"
)
/**
	https://www.geeksforgeeks.org/next-greater-element/
 */
func largestElementOnRight(a []int) []int {

	result := make([]int, len(a))

	for i := 0; i < len(a); i++ {
		result[i] = -1
	}

	stack := stack.Stack{}
	stack.Push(0)

	for i := 1; i < len(a); i++ {
		for !stack.IsEmpty() {
			t, _ := stack.Peek()
			if a[t] < a[i] {
				result[t] = i
				stack.Pop()
			} else {
				break
			}

		}
		stack.Push(i)
	}
	return result
}

func main() {
	arr := []int{8, 10, 2, 4, 20, 17, 39, 1, -2, 4}

	result := largestElementOnRight(arr)

	for i := 0; i < len(result); i++ {
		if result[i] == -1 {
			fmt.Printf("No largest element found for %d\n", arr[i])
		} else {
			fmt.Printf("Largest element for %d  ==> %d \n", arr[i], arr[result[i]])
		}
	}
}

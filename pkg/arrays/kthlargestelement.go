package main

import "fmt"


/**
 * Kth largest element in an array.
 * Use quickselect of quicksort to find the solution in O(nlogn) time.
*/
func findKthLargestElement(arr []int, left , right , k int) int {

	// Check for valid input values.
	if len(arr) < 1 || k < 1 || k > len(arr) || left > right {
		return -1
	}

	// Size of an input array
	size := len(arr)

	// final desired index we're looking for.
	finalIndex := size - k

	// Alternatively, you can modify the above line as follows to get kth min element instead
	// let finalIndex = k - 1

	// Make the partition and find the index of result pivot element.
	partition := partition(arr, left, right)

	// If partition index is same as our expected final index, we've found kth max/min element.
	if partition == finalIndex {
		return arr[partition]
	}

	// If parition falls on the left side of desired index,
	// we will start looking from successive partition index to the end
	if partition < finalIndex {
		return findKthLargestElement(arr, partition + 1, right, k)
	} else {
		// If parition falls on the right side of desired index,
		// we will start looking from beginning to the previous partition index
		return findKthLargestElement(arr, left, partition - 1, k)
	}
}

func partition(arr []int, left , right int) int {

	if left == right {
		return left
	}
	i := left +1
	j := right

	for {
		for arr[i] < arr[left] && i < right {
			i++
		}

		for arr[j] > arr[left] && j > left {
			j--
		}
		if i >= j {
			break
		}

		temp1 :=  arr[i]
		arr[i] = arr[j]
		arr[j] = temp1

		i++
		j--
	}

	temp := arr[left]
	arr[left] = arr[j]
	arr[j] = temp

	return j

}

func main() {
	arr := []int{100, 2, 1000, 3, 4, 5, 6, 10, 11}
	kthLargest := findKthLargestElement(arr, 0,8, 2)
	fmt.Printf("Kth largest element is  : %d\n", kthLargest)
}

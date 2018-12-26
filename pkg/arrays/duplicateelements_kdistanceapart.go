package main

import "fmt"

/**
 * https://www.geeksforgeeks.org/check-given-array-contains-duplicate-elements-within-k-distance/
 */
func duplicate(arr []int, k int) bool {

	visited := make(map[int]struct{}, len(arr))

	for i := 0; i < len(arr); i++ {

		_, ok := visited[arr[i]]

		if ok {
			return true
		}

		if i >= k {
			delete(visited, arr[i-k])
		}

		visited[arr[i]] = struct{}{}
	}
	return false
}

func main() {
	arr := []int{1, 2, 3, 4, 1, 2, 3, 4}
	duplicatesFound := duplicate(arr, 3)
	if duplicatesFound {
		fmt.Println(" found duplicates k distance apart")
	} else {
		fmt.Println(" no duplicates found k distance apart")
	}
}

package main

import "fmt"
/*
John works at a clothing store. He has a large pile of socks that he must pair by color for sale.
Given an array of integers representing the color of each sock, determine how many pairs of socks with matching colors there are.

For example, there are
socks with colors . There is one pair of color and one of color .
There are three odd socks left, one of each color. The number of pairs is .
*/

func findSocksPairs(a []int) int {
	pairCount := 0
	socksCounts := make(map[int]int, 100)
	for i := 0; i < len(a); i++ {
		socksId := a[i]
		sockCount, ok := socksCounts[socksId]
		if !ok {
			sockCount++
		} else if sockCount == 1 {
			sockCount = 0
			pairCount++
		} else {
			sockCount++
		}
		socksCounts[socksId] = sockCount
	}
	return pairCount
}

func main() {
	arr := []int{4, 5, 5, 5, 6, 6, 4, 1, 4, 4, 3, 6, 6, 3, 6, 1, 4, 5, 5, 5}
	socksPairs := findSocksPairs(arr)
	fmt.Printf("Socks Pairs : %d\n", socksPairs)
}

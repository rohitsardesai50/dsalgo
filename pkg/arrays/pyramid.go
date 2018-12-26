package main

import (
	"fmt"
)
/**
 Generate pyramid pattern like below
        1
       2 3
      4 5 6
     7 8 9
 */
func main() {

	n := 9
	k := 1
	for i := 1; i <= n; i++ {
		for space := n-1; space >= i; space-- {
			fmt.Print(" ")

		}

		for j := 1; j <= i; j++ {
			if k > n {
				break
			}
			output := fmt.Sprintf("%d%s", k, " ")
			fmt.Print(output)
			k++
		}

		fmt.Println("")
	}
}
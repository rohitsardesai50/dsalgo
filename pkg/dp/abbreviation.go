package main

import (
	"unicode"
	"fmt"
)

/*You can perform the following operations on the string,


1.Capitalize zero or more of a's lowercase letters. Delete all of the remaining lowercase letters in "a"

2.Delete all of the remaining lowercase letters in "a"

Given two strings, "a" and "b" , determine if it's possible to make equal to as described.
If so, print YES on a new line. Otherwise, print NO.
*/


var possible bool

func canCapitalize(a string, b string)  {

	fmt.Printf("===========Inside canCapitalize============\n")
	fmt.Printf("possible : %v\n", possible)
	if  possible || len(a) < len(b) {
		fmt.Printf("return1\n")
		return
	}

	if len(b) == 0 {
		if checkLowercase(a) || len(a) == 0 {
			fmt.Printf("len(a) : %d\n", len(a))
			possible = true
			fmt.Printf("return2\n")
			return
		}
	}

	fmt.Printf("first string : %s, second string : %s\n", a, b)
	runesA := []rune(a)
	fc := runesA[0]

	runedA := string(runesA[1:])

	if unicode.IsLower(fc) {
		canCapitalize(runedA, b)
	}
	if b != "" {
		runesB := []rune(b)
		if unicode.ToUpper(fc) != runesB[0] {
			fmt.Printf("return3\n")
			return
		}
		runedB := string(runesB[1:])
		canCapitalize(runedA, runedB)
	}



	return
}

func checkLowercase(s string) bool {
	for _, a := range s {
		if unicode.IsLower(a) {
			continue
		} else {
			return false
		}
	}
	return true
}



func main() {

	possible = false
	canCapitalize("beFgH", "EFG")
	if possible {
		fmt.Printf("can capitalize")
	} else{
		fmt.Printf("can't capitalize")
	}

}
package main

import "fmt"

/**
 * https://www.geeksforgeeks.org/find-a-tour-that-visits-all-stations/
 */
func findStartingPointOfTour(availableGasArr []int, requiredGasArr []int) int {

	startPoint := -1
	endPoint := 0
	currGas := 0
	visitedOnce := false

	for startPoint != endPoint {

		currGas += availableGasArr[endPoint] - requiredGasArr[endPoint]

		if startPoint == -1 {
			startPoint = endPoint
		}

		if endPoint == len(availableGasArr)-1 && visitedOnce == false {
			visitedOnce = true
		} else if endPoint == len(availableGasArr)-1 && visitedOnce == true {
			return -1
		}

		if currGas < 0 {
			startPoint = -1
			currGas = 0
		}
		endPoint = (endPoint + 1) % len(availableGasArr)
	}
	return endPoint
}

func main() {
	availableGasArr := []int{4, 4, 7, 4}
	requiredGasArr := []int{6, 5, 3, 5}
	startingPoint := findStartingPointOfTour(availableGasArr, requiredGasArr)
	if startingPoint == -1 {
		fmt.Println(" no gas tour exists which visits all gas stations")
	} else {
		fmt.Printf(" gas tour can be completed if started at station :[%d]", startingPoint)
	}
}

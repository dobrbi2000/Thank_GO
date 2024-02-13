package main

import (
	"fmt"
)

func isPositive(num int) bool {
	return num > 0
}

func filter(predicate func(int) bool, iterable []int) []int {
	var result []int
	for _, num := range iterable {
		if predicate(num) {
			result = append(result, num)
		}
	}
	return result
}

func main() {

	numbers := []int{-1, 2, 0, 5, -9, 10}

	positiveNumbers := filter(isPositive, numbers)

	fmt.Println(positiveNumbers)

}

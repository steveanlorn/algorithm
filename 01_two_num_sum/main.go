package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{3, 5, -4, 8, 11, 1, -1, 6}
	sum := 10
	fmt.Println(twoNumberSum(numbers, sum))
	fmt.Println(twoNumberSumMap(numbers, sum))
	fmt.Println(twoNumberSumSort(numbers, sum))
}

// O(n^2) time | O(1) space
func twoNumberSum(numbers []int, sum int) []int {
	for i, n1 := range numbers {
		n2 := sum - n1

		for j := i + 1; j < len(numbers); j++ {
			if numbers[j] == n2 {
				return []int{n1, n2}
			}
		}
	}

	return nil
}

// O( nlog(n) ) time | O(1) space
func twoNumberSumSort(numbers []int, sum int) []int {
	sort.Ints(numbers)

	left := 0
	right := len(numbers) - 1

	for left < right {
		tempSum := numbers[left] + numbers[right]
		if tempSum == sum {
			return []int{numbers[left], numbers[right]}
		} else if tempSum < sum {
			left += 1
		} else {
			right -= 1
		}
	}

	return nil
}

// O(n) time | O(n) space
func twoNumberSumMap(numbers []int, sum int) []int {
	numbersMap := make(map[int]struct{}, 0)
	for _, n2 := range numbers {
		n1 := sum - n2

		if _, ok := numbersMap[n1]; ok {
			return []int{n1, n2}
		} else {
			numbersMap[n2] = struct{}{}
		}
	}

	return nil
}

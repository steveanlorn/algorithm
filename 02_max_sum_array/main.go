package main

import (
	"fmt"
)

func main() {
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSumArrayWithKadane(arr))
}

func maxSumArrayWithKadane(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	if len(arr) == 1 {
		return arr[1]
	}

	localMax, globalMax := arr[0], arr[0]
	for i := 1; i < len(arr); i++ {
		localMax = max(localMax+arr[i], arr[i])
		globalMax = max(globalMax, localMax)
	}
	return globalMax
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 1, 1, 2, 2, 6, 5, 6}
	occurrences := make(map[int]int)

	for _, num := range nums {
		occurrences[num]++
	}
	fmt.Println(occurrences)
}

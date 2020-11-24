package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(nums))
}

// FUNC_SUM_START OMIT
func sum(ix []int) int {
	total := 0
	for _, i := range ix {
		total += i
	}
	return total
}

// FUNC_SUM_END OMIT

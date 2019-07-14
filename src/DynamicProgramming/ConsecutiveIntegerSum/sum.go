package main

import (
	"fmt"
)

func main() {

	array := [7]int{5, 15, -30, 10, -5, 40, 10}
	conseq := [7]int{}

	max := array[0]
	conseq[0] = max

	for i := 1; i < len(array); i++ {

		sum := conseq[i-1] + array[i]

		if sum > 0 {
			max = sum
			conseq[i] = max
		} else {
			conseq[i] = 0
		}
	}

	fmt.Println(max)
}

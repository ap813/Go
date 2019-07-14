package main

/*
	From: Algorithms by Dasgupta, C. H. Papadimitriou, and U. V. Vazirani
	A contiguous subsequence of a list S is a subsequence made up of consecutive elements of S. For instance, if S is
	5, 15, −30, 10, −5, 40, 10,
	then 15, −30, 10 is a contiguous subsequence but 5, 15, 40 is not. Give a linear-time algorithm for
	the following task:
	Input: A list of numbers, a1, a2, . . . , an.
	Output: The contiguous subsequence of maximum sum (a subsequence of length zero has sum zero).
	For the preceding example, the answer would be 10, −5, 40, 10, with a sum of 55.
*/
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

		if sum > conseq[i] {
			max = sum
			conseq[i] = max
		} else {
			conseq[i] = array[i]
		}
	}

	if max < 0 {
		max = 0
	}

	fmt.Println(max)
}

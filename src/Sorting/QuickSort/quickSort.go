package main

import (
	"fmt"
)

// QuickSort algoirthm
func quickSort(array []int) []int{

	// Base Case: return the array if length is less than 2
	if len(array) < 2 { return array }

	// Left is the start and Right is the end
	left, right := 0, len(array) - 1

	// Pick a Partition
	partition := len(array) / 2;

	// Swap
	array[partition], array[right] = array[right], array[partition]

	// Things smaller than the partition to the swapped left
	for i := range array {
		if array[i] < array[right] {
			// Swap array[i] with array[left]
			array[i], array[left] = array[left], array[i];
			left++;
		}
	}

	// Place the pivot after the last smaller element
	array[left], array[right] = array[right], array[left]

	// Recursively call Quicksort using Slices
	quickSort(array[:left]);
	quickSort(array[left + 1:]);

	// Return Sorted Array
	return array;
}

func main() {

	array := []int{12,91,48,63,1,39,22};

	array = quickSort(array[:]);

	fmt.Println(array);

}

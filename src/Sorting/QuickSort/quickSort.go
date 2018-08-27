package main

import (
	"fmt"
)

// QuickSort algoirthm
func quickSort(a []int) []int{

	// Base Case: return the array if length is less than 2
	if len(a) < 2 { return a }

	// Left is the start and Right is the end
	left, right := 0, len(a) - 1

	// Pick a Partition
	partition := len(a) / 2;

	// Swap
	a[partition], a[right] = a[right], a[partition]

	// Things smaller than the partition to the swapped left
	for i := range a {
		if a[i] < a[right] {
			a[i], a[left] = a[left], a[i];
			left++;
		}
	}

	// Place the pivot after the last smaller element
	a[left], a[right] = a[right], a[left]

	// Recursively call Quicksort using Slices
	quickSort(a[:left]);
	quickSort(a[left + 1:]);

	// Return Sorted Array
	return a;
}

func main() {

	array := []int{12,91,48,63,1,39,22};

	array = quickSort(array[:]);

	fmt.Println(array);

}

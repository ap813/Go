package main

import "fmt"

func insertionSort(a [7]int, length int) [7]int {

	array := a;

	// Sorting
	for i := 1; i < length; i++ {

		for j := 0; j < i; j++ {

			if array[i] < array[j] {

								

			}

		}

	}

	return array;
}

func main() {

	var array = [7]int{1,3,51,6,12,7,91};

	fmt.Println(insertionSort(array, 7));
}

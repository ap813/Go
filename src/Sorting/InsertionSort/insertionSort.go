package main

import "fmt"

func insertionSort(a [7]int, length int) [7]int {

	array := a;

	largest := -1;
	index := 0;

	// Sorting
	for i := 0; i < length-1;  i++{

		for j := 0; j < length - i ; j++ {

			if largest < array[j]  {

				largest = array[j];
				index = j;

			}

		}

		fmt.Println("Largest: ", largest, i);

		temp := array[length-i-1];
		array[length-i-1] = array[index];
		array[index] = temp;

		largest = -1;
		index = 0;

	}

	return array;
}

func main() {

	var array = [7]int{1,3,51,6,12,7,91};

	fmt.Println(insertionSort(array, 7));
}

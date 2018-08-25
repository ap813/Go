package main

import "fmt"

func bubbleSort(a [7]int, length int) [7]int {

	array := a;

	for i := 0; i < length; i++ {
		for j := 0; j < length-1; j++ {
			if array[j] > array[j+1] {
				temp := array[j];
				array[j] = array[j+1];
				array[j+1] = temp;
			}
		}
	}

	return array;
}

func main() {

	a := [7]int{1,41,72,32,69,12,44};

	// Call Bubble Sort
	array := bubbleSort(a, 7);

	fmt.Println(array);
}

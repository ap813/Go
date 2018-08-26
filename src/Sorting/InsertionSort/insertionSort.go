package main

import "fmt"

func insertionSort(a [7]int, length int) [7]int {

	array := a;

	// Sorting
	for i := 1; i < length; i++ {

		for j := 0; j <= i; j++ {

			if (array[i] < array[j]) {

				temp := array[j];
				array[j] = array[i];
				for k := j+1; k <= i ; k++ {
					temp2 := array[k];
					array[k] = temp;
					temp = temp2;
				}

				j = i+1;
			}

		}

	}

	return array;
}

func main() {

	var array = [7]int{1,5,12,8,2,4,5};

	fmt.Println(insertionSort(array, 7));
}

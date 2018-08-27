package main

import "fmt"

func mergeSort(array []int) []int {

	if len(array) < 2 {
		return array;
	}

	middle := len(array) / 2;

	return merge(mergeSort(array[:middle]),mergeSort(array[middle:]));
}

// Create a sorted array
func merge(left []int, right []int) []int {

	length, i, j := len(left) + len(right), 0 ,0;

	// Array will take all the values from left and right
	array := make([]int, length, length);

	// Go thru left and right and insert as you go
	for k := 0; k < length; k++ {
		if j > len(right)-1 && i <= len(left)-1 {
			// Right is finished, but left is not
			array[k] = left[i]
			i++
		} else if i > len(left)-1 && j <= len(right)-1 {
			// Left is finished, but right is not
			array[k] = right[j]
			j++
		} else if left[i] > right[j] {
			// Right is smaller than Left
			array[k] = right[j]
			j++
		} else {
			// Left is smaller than Right
			array[k] = left[i]
			i++
		}
	}

	return array;
}

func main() {
	array := []int{12,91,48,63,1,39,22};

	fmt.Println(mergeSort(array));
}

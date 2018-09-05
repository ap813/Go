package main

import "fmt"

const size int = 100;

type heap struct {
	array [size] int
	length int
}

func createHeap(num int) *heap {
	newHeap := new(heap);
	newHeap.array[0] = num;
	newHeap.length = 1;

	return newHeap;
}

// Insert into a heap
func (h* heap) insert(num int) {
	// Make sure there isn't overflow
	if h.length == size {
		fmt.Println("Heap is full");
	}

	h.array[h.length] = num;

	// Bring larger values up
	for i := h.length; h.array[i] > h.array[i/2];  i=i/2 {
		temp := h.array[i];
		h.array[i] = h.array[i/2];
		h.array[i/2] = temp;
	}

	// Increment length
	h.length++;
}

func main() {
	heap := createHeap(10);

	heap.insert(20);
	heap.insert(40);
	heap.insert(15);
	heap.insert(50);
	heap.insert(12);

	fmt.Println(heap);
}

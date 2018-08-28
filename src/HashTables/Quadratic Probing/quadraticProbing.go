package main

import (
	"fmt"
	"hash/fnv"
)

const INITIAL_SIZE int = 11;


type hash struct {
	table []string
	max int
	size int
}

// Generates a Hash Value from a string
func (h *hash) hash(str string) uint32 {
	hashed := fnv.New32a();
	hashed.Write([]byte(str));
	return hashed.Sum32();
}

// Called when the number of elements
// in the table is larger than half
func (h *hash) IncreaseSize() {

	//oldHash := h.table;
	//oldHashLength := h.max;

	// The first prime number would be
	// twice the size of the original
	// plus 1
	start := (h.max * 2) + 1;

	// Iterate the odd numbers till
	// We find a prime number
	for i := start; i > 0; i=i+2 {
		if isPrime(i) {
			h.max = i;
			i = -2;
		}
	}

	// Figure out how to make array of strings
	// the Size of h.max
	//h.table = [h.max]string;

	//for j := 0; j < oldHashLength; j++  {
	//	if oldHash[j] != "" {
	//
	//	}
	//}

	return;
}

// Check if an Integer is Prime
// Returns Bool
func isPrime(num int) bool{

	if num == 0 {
		return false;
	}

	for i := 2; i < num-1; i++ {
		if num % i == 0 {
			return false;
		}
	}

	return true;
}

func createHashTable() *hash {

	newHash := new(hash);

	newHash.max = INITIAL_SIZE;

	newHash.size = 0;

	return newHash;
}

func main() {

	hashtable := createHashTable();

	hashtable.IncreaseSize()

	fmt.Println(hashtable.max);
}

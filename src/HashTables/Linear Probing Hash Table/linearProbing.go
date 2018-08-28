package main

import (
	"fmt"
	"hash/fnv"
)

const size int = 10;

type hash struct {
	table [size]string
}

func (h *hash) insert(str string) {

	var index int = int(h.hash(str)) % size;

	fmt.Println(index);

	for i := index; i >= 0; i++ {
		if h.table[i % size] == "" || h.table[i % size] == str {
			fmt.Println("Inserting: ", str);
			h.table[i % size] = str;
			i = -2;
		}
	}
}

func (h *hash) search(str string) bool {

	var index int = int(h.hash(str)) % size;

	for i := index; i>=0 ; i++ {
		if h.table[i % size] == "" {
			return false;
		}

		if h.table[i % size] == str {
			return true;
		}
	}

	return false;
}

func (h *hash) hash(str string) uint32 {
	hashed := fnv.New32a();
	hashed.Write([]byte(str));
	return hashed.Sum32();
}

func main() {

	table := new(hash);

	table.insert("Steven");
	table.insert("Mark");
	table.insert("Alex");
	table.insert("Josh");
	table.insert("Kobe");
	table.insert("Sarah");
}

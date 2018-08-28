package main

import "fmt"

type queue struct {
	head *node
	tail *node
	n int
}

type node struct {
	next *node
	back *node
	value int
}

func (q *queue) enqueue(num int) {

	temp := new(node);
	temp.value = num;

	if q.n == 0 {
		q.head = temp;
		q.tail = temp;
		q.n++;
		return;
	}

	newNode := temp;
	newNode.next = q.head;
	q.head.back = newNode;
	q.head = newNode;

	q.n++;

}

func (q *queue) dequeue() {

	if q.n == 0 {
		fmt.Println("Queue is Empty");
		return;
	}

	fmt.Println("Popping: ", q.tail.value);

	if q.n == 1 {
		q.head = nil;
		q.tail = nil;
		q.n = 0;
		return;
	}


	temp := q.tail.back;
	temp.next = nil;
	q.tail = temp;

 	q.n--;
}

func (q *queue) print() {
	for node := q.head ; node != nil; node = node.next {
		fmt.Println(node.value);
	}
}

func (q *queue) length() int {
	return q.n;
}

func createQueue() *queue {

	q := new(queue);

	q.n = 0;

	return q;
}

func main() {

	q := createQueue();

	q.print();

	q.enqueue(10);

	q.enqueue(20);

	q.enqueue(30);

	q.enqueue(40);

	q.dequeue();

	q.dequeue();

	q.dequeue();

	q.dequeue();

	q.print();
}
